package implementations

import (
	"auth-service/lib/adapter/messageBroker"
	"fmt"
	"log/slog"
	"os"
	"time"
)

var (
	DEBUG = 0
	INFO  = 1
	ERROR = 2
)

const AppName = "auth-service"

type Logger struct {
	log    *slog.Logger
	broker messageBroker.MessageBroker
}

func NewLogger(broker messageBroker.MessageBroker, env string) *Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "dev":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "prod":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return &Logger{
		log:    log,
		broker: broker,
	}
}

func (l *Logger) Debug(msg string, args ...any) {
	l.log.Debug(msg, args...)
	l.sendToBroker(DEBUG, msg, args)
}

func (l *Logger) Info(msg string, args ...any) {
	l.log.Info(msg, args...)
	l.sendToBroker(INFO, msg, args)
}

func (l *Logger) Error(msg string, args ...any) {
	l.log.Error(msg, args...)
	l.sendToBroker(ERROR, msg, args)
}

func (l *Logger) sendToBroker(level int, msg string, args ...any) {
	err := l.broker.Publish(messageBroker.LogMessage{
		AppName: AppName,
		Level:   level,
		Time:    time.Now(),
		Text:    fmt.Sprintf(msg, args...),
	})
	if err != nil {
		l.log.Error(fmt.Sprintf("Error publishing message: %s", err))
	}
}
