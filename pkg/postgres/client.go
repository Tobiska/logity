package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"logity/config"
	"time"
)

const (
	Dialect = `postgres`
)

func New(cfg *config.Database) (*pgxpool.Pool, error) {
	pgxCfg, err := pgxpool.ParseConfig(cfg.Dsn)
	if err != nil {
		return nil, fmt.Errorf("pgxpool cfg parse error: %w", err)
	}
	pgxCfg.MaxConns = int32(cfg.MaxIdleConn)
	pgxCfg.MaxConnLifetime = time.Duration(cfg.MaxLifeTimeConn)

	conn, err := pgxpool.NewWithConfig(context.TODO(), pgxCfg)
	if err != nil {
		return nil, fmt.Errorf("pgxpool new error: %w", err)
	}

	if err := conn.Ping(context.TODO()); err != nil {
		return nil, fmt.Errorf("pgxpool ping error: %w", err)
	}

	return conn, nil
}
