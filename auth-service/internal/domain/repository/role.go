package repository

import (
	"auth-service/internal/domain/model"
	"context"
)

type RoleRepository interface {
	GetRoleByName(ctx context.Context, name string) (model.Role, error)
	GetRoleByID(ctx context.Context, id int) (model.Role, error)
}
