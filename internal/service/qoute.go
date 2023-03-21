package service

import (
	"github.com/PickHD/quote-echo-grpc/configs"
	"github.com/PickHD/quote-echo-grpc/internal/repository"
)

type (
	QuoteService interface{}

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
