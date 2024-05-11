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

type RoleRepository struct {
	db db.DBAdapter
}

func NewRoleRepository(db db.DBAdapter) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) GetRoleByName(ctx context.Context, name string) (model.Role, error) {
	query := `SELECT id, role FROM roles WHERE role = $1;`
	var role dbModel.Role
	err := r.db.QueryRow(ctx, &role, query, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Role{}, repository.ErrNotFound
		}
		return model.Role{}, err
	}
	return model.Role(role), nil
}

func (r *RoleRepository) GetRoleByID(ctx context.Context, id int) (model.Role, error) {
	query := `SELECT id, role FROM roles WHERE id = $1;`
	var role dbModel.Role
	err := r.db.QueryRow(ctx, &role, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Role{}, repository.ErrNotFound
		}
		return model.Role{}, err
	}
	return model.Role(role), nil
}
