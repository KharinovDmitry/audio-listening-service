package grpc

import (
	"auth-service/internal/service"
	"auth-service/internal/storage"
	"context"
	"errors"
	"github.com/KharinovDmitry/audio-listening-service/protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService interface {
	Login(ctx context.Context, login string, password string) (token string, err error)
	RegisterNewListener(ctx context.Context, login string, password string) (userID int64, err error)
	RegisterNewArtist(ctx context.Context, login string, password string, name string) (userID int64, err error)
}

type serverAPI struct {
	auth.UnimplementedAuthServer
	auth AuthService
}

func Register(gRPCServer *grpc.Server, authService AuthService) {
	auth.RegisterAuthServer(gRPCServer, &serverAPI{auth: authService})
}

func (s *serverAPI) Login(ctx context.Context, in *auth.LoginRequest) (*auth.LoginResponse, error) {
	if in.Login == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	token, err := s.auth.Login(ctx, in.GetLogin(), in.GetPassword())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, "invalid email or password")
		}
		return nil, status.Error(codes.Internal, "failed to login")
	}

	return &auth.LoginResponse{Token: token}, nil
}

func (s *serverAPI) RegisterListener(ctx context.Context, in *auth.RegisterListenerRequest) (*auth.RegisterResponse, error) {
	if in.Login == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	uid, err := s.auth.RegisterNewListener(ctx, in.GetLogin(), in.GetPassword())
	if err != nil {
		// Ошибку storage.ErrUserExists мы создадим ниже
		if errors.Is(err, storage.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}

		return nil, status.Error(codes.Internal, "failed to register user")
	}

	return &auth.RegisterResponse{UserId: uid}, nil
}

func (s *serverAPI) RegisterArtist(ctx context.Context, in *auth.RegisterArtistRequest) (*auth.RegisterResponse, error) {
	if in.Login == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	uid, err := s.auth.RegisterNewListener(ctx, in.GetLogin(), in.GetPassword())
	if err != nil {
		// Ошибку storage.ErrUserExists мы создадим ниже
		if errors.Is(err, storage.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}

		return nil, status.Error(codes.Internal, "failed to register user")
	}

	return &auth.RegisterResponse{UserId: uid}, nil
}
