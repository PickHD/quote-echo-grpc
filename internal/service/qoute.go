package service

import (
	"github.com/PickHD/quote-echo-grpc/configs"
	"github.com/PickHD/quote-echo-grpc/internal/repository"
)

type (
	QouteService interface{}

	QouteServiceImpl struct {
		Config configs.Config
		QouteRepository repository.QouteRepository
	}
)

func NewQouteRepositoryImpl(cfg configs.Config, qouteRepo repository.QouteRepository) QouteService {
	return &QouteServiceImpl{
		Config: cfg,
		QouteRepository: qouteRepo,
	}
}
