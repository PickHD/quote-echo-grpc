package repository

import (
	"github.com/PickHD/quote-echo-grpc/configs"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type (
	QouteRepository interface{}

	QouteRepositoryImpl struct {
		Config configs.Config
		DB     *pgxpool.Pool
		Logger *zap.Logger
	}
)

func NewQouteRepositoryImpl(cfg configs.Config, db *pgxpool.Pool, logger *zap.Logger) QouteRepository {
	return &QouteRepositoryImpl{
		Config: cfg,
		DB:     db,
		Logger: logger,
	}
}
