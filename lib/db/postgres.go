package db

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type PostgresAdapter struct {
	connection *sqlx.DB
}

func NewPostgresAdapter() *PostgresAdapter {
	return &PostgresAdapter{}
}

func (pa *PostgresAdapter) Connect(ctx context.Context, connectionString string) (*sqlx.DB, error) {
	conn, err := sqlx.ConnectContext(ctx, "postgres", connectionString)
	pa.connection = conn
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (pa *PostgresAdapter) Close() error {
	return pa.connection.Close()
}

func (pa *PostgresAdapter) Execute(requestCtx context.Context, sql string, args ...any) error {
	if pa.connection == nil {
		return errors.New("[ PostgresAdapter ] Execute: connection is nil")
	}

	ctx, cancel := context.WithTimeout(requestCtx, TimeoutDbContext)
	defer cancel()

	_, err := pa.connection.NamedExecContext(ctx, sql, args)
	return err
}

func (pa *PostgresAdapter) ExecuteAndGet(requestCtx context.Context, destination any, sql string, args ...any) error {
	if pa.connection == nil {
		return errors.New("[ PostgresAdapter ] ExecuteAndGet: connection is nil")
	}

	ctx, cancel := context.WithTimeout(requestCtx, TimeoutDbContext)
	defer cancel()

	return pa.connection.GetContext(ctx, destination, sql, args)
}

func (pa *PostgresAdapter) Query(requestCtx context.Context, destination any, query string, args ...any) error {
	if pa.connection == nil {
		return errors.New("[ PostgresAdapter ] Query: connection is nil")
	}

	ctx, cancel := context.WithTimeout(requestCtx, TimeoutDbContext)
	defer cancel()

	return pa.connection.SelectContext(ctx, destination, query, args)
}
