package infra

import (
	"context"
	"fmt"
	"os"

	"github.com/PickHD/quote-echo-grpc/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgreSQLPool(cfg *configs.Config) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cfg.DBUsername,
		cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)

	dbpool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)

		return nil, err
	}

	return dbpool, nil
}
