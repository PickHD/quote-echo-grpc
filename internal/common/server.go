package common

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	serverTimeout = 5 * time.Second
)

type GrpcServer interface {
	Run()
	Register()
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
	logger      Logger
}

func NewServer(name string, router Router, infraCloser InfraCloser, logger Logger) *Server {
	return &Server{name, router, infraCloser, logger}
}

func (s *Server) Serve() {
	s.router.Run()

	done := make(chan bool, 1)
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		ctx, cancel := context.WithTimeout(context.Background(), serverTimeout)
		defer cancel()
		s.GracefulStop(ctx, done)
	}()

	<-done
}

func (s *Server) GracefulStop(ctx context.Context, done chan bool) {
	err := s.router.GracefulStop(ctx)
	if err != nil {
		s.logger.Error("failed GracefulStop", err)
	}

	if err = s.infraCloser.Close(); err != nil {
		s.logger.Error("failed infraCloser.Close()", err)
	}

	s.logger.Info("gracefully shutdowned")
	done <- true
}
