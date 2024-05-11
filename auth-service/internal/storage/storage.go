package storage

import (
	"auth-service/internal/domain/repository"
	"auth-service/internal/storage/implementations"
	"auth-service/lib/adapter/db"
	"auth-service/lib/adapter/db/postgres"
	"context"
	_ "github.com/lib/pq"
	"time"
)

type Storage struct {
	db             db.DBAdapter
	UserRepository repository.UserRepository
	RoleRepository repository.RoleRepository
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Init(timeout time.Duration, connStr string) error {
	adapter := postgres.NewPostgresAdapter(timeout)
	_, err := adapter.Connect(context.Background(), connStr)
	if err != nil {
		return err
	}

	s.db = adapter

	s.UserRepository = implementations.NewUserRepository(s.db)
	s.RoleRepository = implementations.NewRoleRepository(s.db)
	return nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
