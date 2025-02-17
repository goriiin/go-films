package auth

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

type DB interface {
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

type Repository struct {
	db  DB
	log *slog.Logger
}
