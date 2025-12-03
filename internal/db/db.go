package db

import (
	"context"
	"time"

	"github.com/AyushDubey63/go-monitor/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(ctx context.Context) (*pgxpool.Pool, error) {
	// Load configuration
	cfgData := config.LoadConfig()
	dsn := cfgData.DatabaseUrl

	// Parse pgx config
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	// Connection pool settings
	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnIdleTime = time.Hour

	// Create connection pool
	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
