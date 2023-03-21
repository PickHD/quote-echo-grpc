package controller

import (
	"context"

	"github.com/PickHD/quote-echo-grpc/internal/common"
)

type Router struct {
	grpcServer common.GrpcServer
}

func NewRouter(grpcServer common.GrpcServer) *Router {
	return &Router{grpcServer}
}

func (r *Router) Run() {
	r.grpcServer.Register()
	r.grpcServer.Run()
}

func (r *Router) GracefulStop(_ context.Context) error {
	return r.grpcServer.GracefulStop()
}
