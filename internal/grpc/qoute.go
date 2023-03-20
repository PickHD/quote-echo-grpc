package grpc

import (
	"github.com/PickHD/quote-echo-grpc/configs"
	"github.com/PickHD/quote-echo-grpc/internal/service"
)

type (
	QouteController interface{}

	QouteControllerImpl struct {
		Config configs.Config
		QouteService service.QouteService
	}
)

func NewQouteControllerImpl(cfg configs.Config, qouteSvc service.QouteService) QouteController {
	return &QouteControllerImpl{
		Config: cfg,
		QouteService: qouteSvc,
	}
}
