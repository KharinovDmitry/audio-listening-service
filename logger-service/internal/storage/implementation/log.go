package implementation

import (
	"context"
	"logger-service/internal/domain/model"
	"logger-service/lib/adapter/db"
	"time"
)

type LogRepository struct {
	db db.DBAdapter
}

func NewLogRepository(db db.DBAdapter) *LogRepository {
	return &LogRepository{
		db: db,
	}
}

func (l *LogRepository) AddLog(ctx context.Context, appName string, level int, time time.Time, message string) error {
	query := `INSERT INTO logs(app_name, level, time, message) VALUES ($1, $2, $3, $4)`

	err := l.db.Execute(ctx, query, appName, level, time, message)
	if err != nil {
		return err
	}
	return nil
}

func (l *LogRepository) GetAllLogs(ctx context.Context) ([]model.LogEvent, error) {
	panic("implement me")
}
