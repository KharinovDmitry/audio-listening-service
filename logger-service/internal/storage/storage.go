package storage

import (
	"context"
	_ "github.com/lib/pq"
	"logger-service/internal/domain/repository"
	"logger-service/internal/storage/implementation"
	"logger-service/lib/adapter/db"
	"logger-service/lib/adapter/db/postgres"
	"time"
)

type Storage struct {
	db            db.DBAdapter
	LogRepository repository.LogRepository
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
	s.LogRepository = implementation.NewLogRepository(adapter)

	return nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
