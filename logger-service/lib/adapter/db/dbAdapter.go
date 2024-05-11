package db

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type DBAdapter interface {
	GetConnection() *sqlx.DB
	Connect(ctx context.Context, connectionString string) (*sqlx.DB, error)
	Close() error
	Execute(requestCtx context.Context, sql string, args ...interface{}) error
	ExecuteAndGet(requestCtx context.Context, destination interface{}, sql string, args ...interface{}) error
	Query(requestCtx context.Context, destination interface{}, query string, args ...interface{}) error
	QueryRow(requestCtx context.Context, destination interface{}, query string, args ...interface{}) error
}
