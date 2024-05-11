package implementations

import (
	"auth-service/internal/domain/model"
	"auth-service/internal/domain/repository"
	"auth-service/internal/domain/service"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

type AuthService struct {
	userRepo     repository.UserRepository
	roleRepo     repository.RoleRepository
	tokenService service.TokenService
	salt         string
}

func NewAuthService(userRepo repository.UserRepository, roleRepo repository.RoleRepository, tokenService service.TokenService, salt string) *AuthService {
	return &AuthService{
		userRepo:     userRepo,
		roleRepo:     roleRepo,
		tokenService: tokenService,
		salt:         salt,
	}
}

func (a *AuthService) SignUp(ctx context.Context, login string, password string, role string) (err error) {
	_, err = a.userRepo.GetUserByLogin(ctx, login)
	if err != nil {
		if !errors.Is(err, repository.ErrNotFound) {
			return err
		}
	} else {
		return service.ErrUserAlreadyExists
	}

	dbRole, err := a.roleRepo.GetRoleByName(ctx, role)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return service.ErrUnknownRole
		}
		return err
	}

	err = a.userRepo.AddUser(ctx, login, a.hash(password), dbRole.ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthService) SignIn(ctx context.Context, login string, password string) (token string, err error) {
	user, err := a.userRepo.GetUserByLogin(ctx, login)
	if err != nil {
		return "", service.ErrInvalidCredentials
	}
	if user.Password != a.hash(password) {
		return "", service.ErrInvalidCredentials
	}

	token, err = a.tokenService.CreateToken([]model.Claim{
		{
			Title: model.IDClaimTitle,
			Value: user.ID,
		},
		{
			Title: model.LoginClaimTitle,
			Value: user.Login,
		},
		{
			Title: model.RoleClaimTitle,
			Value: user.Role.Name,
		},
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *AuthService) hash(item string) string {
	encoder := sha256.New()
	encoder.Write([]byte(item + a.salt))
	return hex.EncodeToString(encoder.Sum(nil))
}
