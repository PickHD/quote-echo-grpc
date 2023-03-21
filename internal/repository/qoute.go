package repository

import (
	"github.com/PickHD/quote-echo-grpc/configs"
	"github.com/PickHD/quote-echo-grpc/internal/common"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	QuoteRepository interface{}

	QuoteRepositoryImpl struct {
		Config *configs.Config
		DB     *pgxpool.Pool
		Logger common.Logger
	}
)

func NewQouteRepositoryImpl(cfg *configs.Config, db *pgxpool.Pool, logger common.Logger) *QuoteRepositoryImpl {
	return &QuoteRepositoryImpl{
		Config: cfg,
		DB:     db,
		Logger: logger,
	}
}
