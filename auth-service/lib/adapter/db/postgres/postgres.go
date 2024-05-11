package postgres

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	PostgresDriverName = "postgres"
)

type PostgresAdapter struct {
	Timeout    time.Duration
	connection *sqlx.DB
}

func NewPostgresAdapter(timeout time.Duration) *PostgresAdapter {
	return &PostgresAdapter{
		Timeout: timeout,
	}
}

func (pa *PostgresAdapter) GetConnection() *sqlx.DB {
	return pa.connection
}

func (pa *PostgresAdapter) Connect(ctx context.Context, connectionString string) (*sqlx.DB, error) {
	conn, err := sqlx.ConnectContext(ctx, PostgresDriverName, connectionString)
	pa.connection = conn
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (pa *PostgresAdapter) Close() error {
	return pa.connection.Close()
}

func (pa *PostgresAdapter) Execute(requestCtx context.Context, sql string, args ...interface{}) error {
	if pa.connection == nil {
		return errors.New("[ PostgresAdapter ] Execute: connection is nil")
	}

	ctx, cancel := context.WithTimeout(requestCtx, pa.Timeout)
	defer cancel()

	_, err := pa.connection.ExecContext(ctx, sql, args...)
	return err
}

func (pa *PostgresAdapter) ExecuteAndGet(requestCtx context.Context, destination interface{}, sql string, args ...interface{}) error {
	if pa.connection == nil {
		return errors.New("[ PostgresAdapter ] ExecuteAndGet: connection is nil")
	}

	ctx, cancel := context.WithTimeout(requestCtx, pa.Timeout)
	defer cancel()

	return pa.connection.GetContext(ctx, destination, sql, args...)
}

func (pa *PostgresAdapter) Query(requestCtx context.Context, destination interface{}, query string, args ...interface{}) error {
	if pa.connection == nil {
		return errors.New("[ PostgresAdapter ] Query: connection is nil")
	}

	ctx, cancel := context.WithTimeout(requestCtx, pa.Timeout)
	defer cancel()

	return pa.connection.SelectContext(ctx, destination, query, args...)
}

func (pa *PostgresAdapter) QueryRow(requestCtx context.Context, destination interface{}, query string, args ...interface{}) error {
	if pa.connection == nil {
		return errors.New("[ PostgresAdapter ] QueryRow: connection is nil")
	}

	ctx, cancel := context.WithTimeout(requestCtx, pa.Timeout)
	defer cancel()

	return pa.connection.QueryRowxContext(ctx, query, args...).StructScan(destination)
}
