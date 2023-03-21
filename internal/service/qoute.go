package service

import (
	"context"

	"github.com/PickHD/quote-echo-grpc/configs"
	"github.com/PickHD/quote-echo-grpc/internal/domain"
	"github.com/PickHD/quote-echo-grpc/internal/repository"
	"github.com/google/uuid"
)

type (
	QuoteService interface {
		CreateQuote(ctx context.Context, req domain.CreateOrUpdateQuoteRequest) (*domain.QuoteResponse, error)
		ListQuote(ctx context.Context) ([]domain.QuoteResponse, error)
		DetailQoute(ctx context.Context, id string) (*domain.QuoteResponse, error)
		UpdateQuote(ctx context.Context, id string, req domain.CreateOrUpdateQuoteRequest) (*domain.QuoteResponse, error)
		DeleteQuote(ctx context.Context, id string) (*domain.QuoteResponse, error)
	}

	QuoteServiceImpl struct {
		Config          *configs.Config
		QuoteRepository repository.QuoteRepository
	}
)

func NewQouteServiceImpl(cfg *configs.Config, qouteRepo repository.QuoteRepository) *QuoteServiceImpl {
	return &QuoteServiceImpl{
		Config:          cfg,
		QuoteRepository: qouteRepo,
	}
}

func (qs *QuoteServiceImpl) CreateQuote(
	ctx context.Context, req domain.CreateOrUpdateQuoteRequest,
) (*domain.QuoteResponse, error) {
	data, err := qs.QuoteRepository.Create(ctx, domain.NewQuote(req.Body))
	if err != nil {
		return nil, err
	}

	return &domain.QuoteResponse{
		ID: data.ID,
	}, nil
}

func (qs *QuoteServiceImpl) ListQuote(ctx context.Context) ([]domain.QuoteResponse, error) {
	items, err := qs.QuoteRepository.List(ctx)
	if err != nil {
		return nil, err
	}

	if len(items) < 1 {
		return []domain.QuoteResponse{}, nil
	}

	quotes := make([]domain.QuoteResponse, len(items))

	for i, q := range items {
		q := q
		quotes[i].ID = q.ID
		quotes[i].Body = q.Body
		quotes[i].CreatedAt = q.CreatedAt
		quotes[i].UpdatedAt = q.UpdatedAt
	}

	return quotes, nil
}

func (qs *QuoteServiceImpl) DetailQoute(ctx context.Context, id string) (*domain.QuoteResponse, error) {
	parseID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	item, err := qs.QuoteRepository.Detail(ctx, parseID)
	if err != nil {
		return nil, err
	}

	quote := domain.QuoteResponse{
		ID:        item.ID,
		Body:      item.Body,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}

	return &quote, nil
}

func (qs *QuoteServiceImpl) UpdateQuote(
	ctx context.Context, id string, req domain.CreateOrUpdateQuoteRequest,
) (*domain.QuoteResponse, error) {
	parseID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	_, err = qs.QuoteRepository.Detail(ctx, parseID)
	if err != nil {
		return nil, err
	}

	data, err := qs.QuoteRepository.Update(ctx, parseID, domain.NewQuote(req.Body))
	if err != nil {
		return nil, err
	}

	return &domain.QuoteResponse{
		ID: data.ID,
	}, nil
}

func (qs *QuoteServiceImpl) DeleteQuote(ctx context.Context, id string) (*domain.QuoteResponse, error) {
	parseID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	_, err = qs.QuoteRepository.Detail(ctx, parseID)
	if err != nil {
		return nil, err
	}

	data, err := qs.QuoteRepository.Delete(ctx, parseID)
	if err != nil {
		return nil, err
	}

	return &domain.QuoteResponse{
		ID: data.ID,
	}, nil
}
