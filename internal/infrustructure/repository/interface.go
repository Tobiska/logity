package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type (
	Client interface {
		Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
		QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
		Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
		Begin(ctx context.Context) (pgx.Tx, error)
		Close()
	}
	HashGenerator interface {
		Hash(ctx context.Context, raw string) (string, error)
		Check(ctx context.Context, raw, hash string) error
	}
)
