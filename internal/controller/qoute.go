package controller

import (
	"context"

	"github.com/PickHD/quote-echo-grpc/configs"
	"github.com/PickHD/quote-echo-grpc/internal/service"
	quotepb "github.com/PickHD/quote-echo-grpc/pkg/api/proto/quote"
)

type (
	QuoteController interface {
		CreateQuote(ctx context.Context, req *quotepb.CreateQuoteRequest) (*quotepb.CreateQuoteResponse, error)
		ListQuote(ctx context.Context, req *quotepb.ListQuoteRequest) (*quotepb.ListQuoteResponse, error)
		DetailQuote(ctx context.Context, req *quotepb.DetailQuoteRequest) (*quotepb.DetailQuoteResponse, error)
		UpdateQuote(ctx context.Context, req *quotepb.UpdateQuoteRequest) (*quotepb.UpdateQuoteResponse, error)
		DeleteQuote(ctx context.Context, req *quotepb.DeleteQuoteRequest) (*quotepb.DeleteQuoteResponse, error)
	}

	QuoteControllerImpl struct {
		Config       *configs.Config
		QuoteService service.QuoteService
		quotepb.UnimplementedQuoteServiceServer
	}
)

func NewQuoteControllerImpl(cfg *configs.Config, quoteSvc service.QuoteService) *QuoteControllerImpl {
	return &QuoteControllerImpl{
		Config:       cfg,
		QuoteService: quoteSvc,
	}
}

func (qc *QuoteControllerImpl) CreateQuote(
	ctx context.Context, req *quotepb.CreateQuoteRequest,
) (*quotepb.CreateQuoteResponse, error) {
	return nil, nil
}

func (qc *QuoteControllerImpl) ListQuote(
	ctx context.Context, req *quotepb.ListQuoteRequest,
) (*quotepb.ListQuoteResponse, error) {
	return nil, nil
}

func (qc *QuoteControllerImpl) DetailQuote(
	ctx context.Context, req *quotepb.DetailQuoteRequest,
) (*quotepb.DetailQuoteResponse, error) {
	return nil, nil
}

func (qc *QuoteControllerImpl) UpdateQuote(
	ctx context.Context, req *quotepb.UpdateQuoteRequest,
) (*quotepb.UpdateQuoteResponse, error) {
	return nil, nil
}

func (qc *QuoteControllerImpl) DeleteQuote(
	ctx context.Context, req *quotepb.DeleteQuoteRequest,
) (*quotepb.DeleteQuoteResponse, error) {
	return nil, nil
}
