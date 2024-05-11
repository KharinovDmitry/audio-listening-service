package service

import (
	"auth-service/internal/domain/service"
	"auth-service/internal/service/implementations"
	"auth-service/internal/storage"
	"auth-service/lib/adapter/messageBroker"
	"time"
)

type Manager struct {
	Auth   service.AuthService
	Logger service.LoggerService
	Token  service.TokenService
}

func NewManager(storage *storage.Storage, broker messageBroker.MessageBroker, env, jwtSecret, salt string, ttl time.Duration) *Manager {
	logger := implementations.NewLogger(broker, env)
	token := implementations.NewToken(jwtSecret, ttl)
	auth := implementations.NewAuthService(storage.UserRepository, storage.RoleRepository, token, salt)
	return &Manager{
		Auth:   auth,
		Logger: logger,
		Token:  token,
	}
}
