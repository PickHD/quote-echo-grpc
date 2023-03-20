package grpc

import (
	"fmt"
	"net"

	"github.com/PickHD/quote-echo-grpc/configs"
	"github.com/PickHD/quote-echo-grpc/internal/common"
	"github.com/PickHD/quote-echo-grpc/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	grpcPort int
	logger   *zap.Logger
	s *grpc.Server
	qouteService service.QouteService
}

func NewGrpcServer(config configs.Config,logger *zap.Logger,qouteService service.QouteService) common.GrpcServer {
	srv := &GrpcServer{
		grpcPort: config.ServerPort,
		logger:   logger,
		s: initializeGrpcServer(),
		qouteService: qouteService,
	}

	return srv
}

func (srv *GrpcServer) Register() {
}

func (srv *GrpcServer) Run() {
	go func() {
		addr := fmt.Sprintf("0.0.0.0:%d",srv.grpcPort)
		
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			srv.logger.Error("cannot listen tcp grpc",zap.Error(err))
		}
		
		srv.logger.Info("Listening and serving GRPC server", zap.String("address", lis.Addr().String()))

		if err := srv.s.Serve(lis); err != nil {
			srv.logger.Error("cannot serve grpc server",zap.Error(err))
		}
	}()
}

func (srv *GrpcServer) GracefulStop() error {
	srv.s.GracefulStop()
	return nil
}

func initializeGrpcServer() *grpc.Server {
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	
	return grpcServer
}