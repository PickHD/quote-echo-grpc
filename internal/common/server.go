package common

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

type GrpcServer interface {
	Run()
	GracefulStop() error
}

type Router interface {
	Run()
	GracefulStop(ctx context.Context) error
}

type InfraCloser interface {
	Close() error
}

type Server struct {
	name        string
	router      Router
	infraCloser InfraCloser
	logger      *zap.Logger
}

func NewServer(name string, router Router, infraCloser InfraCloser, logger *zap.Logger) *Server {
	return &Server{name, router, infraCloser, logger}
}

func (s *Server) Serve() {
	s.router.Run()

	done := make(chan bool, 1)
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		s.GracefulStop(ctx, done)
	}()

	<-done
}

func (s *Server) GracefulStop(ctx context.Context, done chan bool) {
	err := s.router.GracefulStop(ctx)
	if err != nil {
		log.Error(err)
	}

	if err = s.infraCloser.Close(); err != nil {
		log.Error(err)
	}

	log.Info("gracefully shutdowned")
	done <- true
}
