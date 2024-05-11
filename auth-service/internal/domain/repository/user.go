package repository

import (
	"auth-service/internal/domain/model"
	"context"
)

type UserRepository interface {
	AddUser(ctx context.Context, login, password string, roleId int) error
	GetUserByLogin(ctx context.Context, login string) (model.User, error)
	GetUserByID(ctx context.Context, id int) (model.User, error)
	UpdateUser(ctx context.Context, id int, newLogin, newPassword string) error
	DeleteUser(ctx context.Context, id int) error
}
