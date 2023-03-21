package repository

import (
	"context"
	"time"

	"github.com/PickHD/quote-echo-grpc/configs"
	"github.com/PickHD/quote-echo-grpc/internal/common"
	"github.com/PickHD/quote-echo-grpc/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	QuoteRepository interface {
		Create(ctx context.Context, req domain.Quote) (*domain.Quote, error)
		List(ctx context.Context) ([]domain.Quote, error)
		Detail(ctx context.Context, id uuid.UUID) (*domain.Quote, error)
		Update(ctx context.Context, id uuid.UUID, req domain.Quote) (*domain.Quote, error)
		Delete(ctx context.Context, id uuid.UUID) (*domain.Quote, error)
	}

	QuoteRepositoryImpl struct {
		Config *configs.Config
		DB     *pgxpool.Pool
		Logger common.Logger
	}
)

func NewQuoteRepositoryImpl(cfg *configs.Config, db *pgxpool.Pool, logger common.Logger) *QuoteRepositoryImpl {
	return &QuoteRepositoryImpl{
		Config: cfg,
		DB:     db,
		Logger: logger,
	}
}

func (qr *QuoteRepositoryImpl) Create(ctx context.Context, req domain.Quote) (*domain.Quote, error) {
	q := `INSERT INTO quotes (body) VALUES ($1) RETURNING id`

	row := qr.DB.QueryRow(ctx, q, req.Body)
	if err := row.Scan(&req.ID); err != nil {
		qr.Logger.Error("QuoteRepositoryImpl.Create Scan Error", err)

		return nil, err
	}

	return &req, nil
}

func (qr *QuoteRepositoryImpl) List(ctx context.Context) ([]domain.Quote, error) {
	q := `
		SELECT
			id,
			body,
			created_at,
			updated_at
		FROM
			quotes
	`

	rows, err := qr.DB.Query(ctx, q)
	if err != nil {
		qr.Logger.Error("QuoteRepositoryImpl.List Query Error", err)

		return nil, err
	}

	var quotes []domain.Quote

	for rows.Next() {
		quote := domain.Quote{}

		err := rows.Scan(
			&quote.ID,
			&quote.Body,
			&quote.CreatedAt,
			&quote.UpdatedAt,
		)
		if err != nil {
			qr.Logger.Error("QuoteRepositoryImpl.List Rows Scan Error", err)

			return nil, err
		}

		quotes = append(quotes, quote)
	}

	return quotes, nil
}

func (qr *QuoteRepositoryImpl) Detail(ctx context.Context, id uuid.UUID) (*domain.Quote, error) {
	q := `
		SELECT
			id,
			body,
			created_at,
			updated_at
		FROM
			quotes
		WHERE
			id = $1
	`
	quote := domain.Quote{}

	row := qr.DB.QueryRow(ctx, q, id)
	err := row.Scan(
		&quote.ID,
		&quote.Body,
		&quote.CreatedAt,
		&quote.UpdatedAt,
	)

	if err != nil {
		qr.Logger.Error("QuoteRepositoryImpl.Detail Row Scan Error", err)

		return nil, err
	}

	return &quote, nil
}

func (qr *QuoteRepositoryImpl) Update(ctx context.Context, id uuid.UUID, req domain.Quote) (*domain.Quote, error) {
	q := `UPDATE quotes SET body = $1, updated_at = $2 WHERE id = $3`

	_, err := qr.DB.Exec(ctx, q, req.Body, time.Now(), id)
	if err != nil {
		qr.Logger.Error("QuoteRepositoryImpl.Update Exec Error", err)

		return nil, err
	}

	req.ID = id

	return &req, nil
}

func (qr *QuoteRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) (*domain.Quote, error) {
	q := `DELETE FROM quotes WHERE id = $1`

	_, err := qr.DB.Exec(ctx, q, id)
	if err != nil {
		qr.Logger.Error("QuoteRepositoryImpl.Delete Exec Error", err)

		return nil, err
	}

	quote := domain.Quote{
		ID: id,
	}

	return &quote, nil
}
