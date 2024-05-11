package implementations

import (
	"auth-service/internal/domain/model"
	"auth-service/internal/domain/repository"
	"auth-service/internal/storage/dbModel"
	"auth-service/lib/adapter/db"
	"context"
	"database/sql"
	"errors"
)

type UserRepository struct {
	db db.DBAdapter
}

func NewUserRepository(db db.DBAdapter) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) AddUser(ctx context.Context, login, password string, roleId int) error {
	query := `INSERT INTO users(login, password, role_id) VALUES ($1, $2, $3)`

	err := u.db.Execute(ctx, query, login, password, roleId)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetUserByLogin(ctx context.Context, login string) (model.User, error) {
	query := `SELECT id, login, password, role_id FROM users WHERE login = $1;`

	var user dbModel.User
	err := u.db.QueryRow(ctx, &user, query, login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, repository.ErrNotFound
		}
		return model.User{}, err
	}

	query = `SELECT id, role FROM roles WHERE id = $1;`
	var role dbModel.Role
	err = u.db.QueryRow(ctx, &role, query, user.RoleID)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:       user.ID,
		Login:    user.Login,
		Password: user.Password,
		Role: model.Role{
			ID:   role.ID,
			Name: role.Name,
		},
	}, nil
}

func (u *UserRepository) GetUserByID(ctx context.Context, id int) (model.User, error) {
	query := `SELECT id, login, password, role_id FROM users WHERE id = $1;`

	var user dbModel.User
	err := u.db.QueryRow(ctx, &user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, repository.ErrNotFound
		}
		return model.User{}, err
	}

	query = `SELECT id, role FROM roles WHERE id = $1;`
	var role dbModel.Role
	err = u.db.QueryRow(ctx, &role, query, user.RoleID)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:       user.ID,
		Login:    user.Login,
		Password: user.Password,
		Role: model.Role{
			ID:   role.ID,
			Name: role.Name,
		},
	}, nil
}

func (u *UserRepository) UpdateUser(ctx context.Context, id int, newLogin, newPassword string) error {
	query := `UPDATE users SET login = $1, password = $2 WHERE id = $3;`

	err := u.db.Execute(ctx, query, newLogin, newPassword, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repository.ErrNotFound
		}
		return err
	}
	return nil
}

func (u *UserRepository) DeleteUser(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1;`

	err := u.db.Execute(ctx, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repository.ErrNotFound
		}
		return err
	}
	return nil
}
