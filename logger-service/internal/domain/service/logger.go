package service

import (
	"context"
	"errors"
)

var (
	ErrLogChanClosed = errors.New("log channel closed")
)

type Logger interface {
	StartWriteLogs(ctx context.Context) error
}
