package server

import (
	"github.com/chuangxinyuan/fengchu/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	logger log.Logger,
	config *config.Server,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			tracing.Server(),
			metadata.Server(),
		),
	}
	if config.Grpc.Network != "" {
		opts = append(opts, grpc.Network(config.Grpc.Network))
	}
	if config.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(config.Grpc.Addr))
	}
	if config.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(config.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	return srv
}
