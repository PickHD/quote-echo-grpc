package controller

import (
	"fmt"
	"net"

	"github.com/PickHD/quote-echo-grpc/configs"
	"github.com/PickHD/quote-echo-grpc/internal/common"
	"github.com/PickHD/quote-echo-grpc/internal/service"
	quotepb "github.com/PickHD/quote-echo-grpc/pkg/api/proto/quote"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	grpcPort        int
	logger          common.Logger
	s               *grpc.Server
	quoteController *QuoteControllerImpl
}

func NewGrpcServer(config *configs.Config, logger common.Logger, quoteService service.QuoteService) *GrpcServer {
	srv := &GrpcServer{
		grpcPort:        config.ServerPort,
		logger:          logger,
		s:               initializeGrpcServer(),
		quoteController: NewQuoteControllerImpl(config, quoteService),
	}

	return srv
}

func (srv *GrpcServer) Register() {
	reflection.Register(srv.s)

	quotepb.RegisterQuoteServiceServer(srv.s, srv.quoteController)
}

func (srv *GrpcServer) Run() {
	go func() {
		addr := fmt.Sprintf("0.0.0.0:%d", srv.grpcPort)

		lis, err := net.Listen("tcp", addr)
		if err != nil {
			srv.logger.Error("cannot listen tcp grpc", err)
		}

		srv.logger.Info("Listening and serving GRPC server", lis.Addr().String())

		if err := srv.s.Serve(lis); err != nil {
			srv.logger.Error("cannot serve grpc server", err)
		}
	}()
}

func (srv *GrpcServer) GracefulStop() error {
	srv.s.GracefulStop()
	return nil
}

func initializeGrpcServer() *grpc.Server {
	return grpc.NewServer()
}
