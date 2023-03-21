package controller

import (
	"context"

	"github.com/PickHD/quote-echo-grpc/configs"
	"github.com/PickHD/quote-echo-grpc/internal/domain"
	"github.com/PickHD/quote-echo-grpc/internal/service"
	quotepb "github.com/PickHD/quote-echo-grpc/pkg/api/proto/quote"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	newQuote := domain.CreateOrUpdateQuoteRequest{
		Body: req.Body,
	}

	data, err := qc.QuoteService.CreateQuote(ctx, newQuote)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed Create Quote %s", err.Error())
	}

	return &quotepb.CreateQuoteResponse{
		Id: data.ID.String(),
	}, nil
}

func (qc *QuoteControllerImpl) ListQuote(
	ctx context.Context, _ *quotepb.ListQuoteRequest,
) (*quotepb.ListQuoteResponse, error) {
	data, err := qc.QuoteService.ListQuote(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed List Quote %s", err.Error())
	}

	if len(data) < 1 {
		return &quotepb.ListQuoteResponse{}, nil
	}

	quotes := make([]*quotepb.Quote, len(data))

	for i, q := range data {
		quotes[i] = &quotepb.Quote{
			Id:   q.ID.String(),
			Body: q.Body,
		}
	}

	return &quotepb.ListQuoteResponse{
		Quotes: quotes,
	}, nil
}

func (qc *QuoteControllerImpl) DetailQuote(
	ctx context.Context, req *quotepb.DetailQuoteRequest,
) (*quotepb.DetailQuoteResponse, error) {
	data, err := qc.QuoteService.DetailQoute(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed List Quote %s", err.Error())
	}

	return &quotepb.DetailQuoteResponse{
		Quote: &quotepb.Quote{
			Id:   data.ID.String(),
			Body: data.Body,
		},
	}, nil
}

func (qc *QuoteControllerImpl) UpdateQuote(
	ctx context.Context, req *quotepb.UpdateQuoteRequest,
) (*quotepb.UpdateQuoteResponse, error) {
	newQuote := domain.CreateOrUpdateQuoteRequest{
		Body: req.GetBody(),
	}

	data, err := qc.QuoteService.UpdateQuote(ctx, req.GetId(), newQuote)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed Update Quote %s", err.Error())
	}

	return &quotepb.UpdateQuoteResponse{
		Id: data.ID.String(),
	}, nil
}

func (qc *QuoteControllerImpl) DeleteQuote(
	ctx context.Context, req *quotepb.DeleteQuoteRequest,
) (*quotepb.DeleteQuoteResponse, error) {
	data, err := qc.QuoteService.DeleteQuote(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed Delete Quote %s", err.Error())
	}

	return &quotepb.DeleteQuoteResponse{
		Id: data.ID.String(),
	}, nil
}
