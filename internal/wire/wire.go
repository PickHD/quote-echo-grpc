package wire

import (
	"github.com/PickHD/quote-echo-grpc/configs"
	"github.com/PickHD/quote-echo-grpc/internal/closer"
	"github.com/PickHD/quote-echo-grpc/internal/common"
	"github.com/PickHD/quote-echo-grpc/internal/controller"
	"github.com/PickHD/quote-echo-grpc/internal/infra"
	"github.com/PickHD/quote-echo-grpc/internal/repository"
	"github.com/PickHD/quote-echo-grpc/internal/service"
	"github.com/google/wire"
)

func InitializeServer(name string) (*common.Server, error) {
	wire.Build(
		configs.NewConfig,

		common.NewLogger,

		infra.NewPostgreSQLPool,

		repository.NewQouteRepositoryImpl,
		wire.Bind(new(repository.QuoteRepository), new(*repository.QuoteRepositoryImpl)),

		service.NewQouteServiceImpl,
		wire.Bind(new(service.QuoteService), new(*service.QuoteServiceImpl)),

		controller.NewQuoteControllerImpl,
		wire.Bind(new(controller.QuoteController), new(*controller.QuoteControllerImpl)),

		controller.NewGrpcServer,
		wire.Bind(new(common.GrpcServer), new(*controller.GrpcServer)),

		controller.NewRouter,
		wire.Bind(new(common.Router), new(*controller.Router)),

		closer.NewInfraCloser,
		wire.Bind(new(common.InfraCloser), new(*closer.InfraCloser)),

		common.NewServer,
	)

	return &common.Server{}, nil
}
