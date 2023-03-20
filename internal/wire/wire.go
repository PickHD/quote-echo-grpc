package wire

import (
	"github.com/PickHD/quote-echo-grpc/internal/common"
	"github.com/google/wire"
)

func InitializeServer(name string) (*common.Server, error) {
	wire.Build()
	return &common.Server{}, nil
}
