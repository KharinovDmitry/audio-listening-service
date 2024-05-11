package implementation

import (
	"context"
	"logger-service/internal/domain/repository"
	"logger-service/internal/domain/service"
	"logger-service/lib/adapter/messageBroker"
)

type Logger struct {
	broker     messageBroker.MessageBroker
	repository repository.LogRepository
}

func NewLogger(repository repository.LogRepository, broker messageBroker.MessageBroker) *Logger {
	return &Logger{
		broker:     broker,
		repository: repository,
	}
}

func (l *Logger) StartWriteLogs(ctx context.Context) error {
	logChan, err := l.broker.Subscribe(messageBroker.LogsQueueName)
	if err != nil {
		return err
	}
	for {
		event, ok := <-logChan
		if !ok {
			return service.ErrLogChanClosed
		}

		err = l.repository.AddLog(ctx, event.AppName, event.Level, event.Time, event.Text)
		if err != nil {
			return err
		}
	}
}
