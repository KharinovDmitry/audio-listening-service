package service

import (
	"logger-service/internal/domain/service"
	"logger-service/internal/service/implementation"
	"logger-service/internal/storage"
	"logger-service/lib/adapter/messageBroker"
)

type Manager struct {
	Logger service.Logger
}

func NewManager(storage *storage.Storage, broker messageBroker.MessageBroker) *Manager {
	logger := implementation.NewLogger(storage.LogRepository, broker)
	return &Manager{
		Logger: logger,
	}
}
