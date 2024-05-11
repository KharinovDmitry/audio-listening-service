package repository

import (
	"context"
	"logger-service/internal/domain/model"
	"time"
)

type LogRepository interface {
	AddLog(ctx context.Context, appName string, level int, time time.Time, message string) error
	GetAllLogs(ctx context.Context) ([]model.LogEvent, error)
}
