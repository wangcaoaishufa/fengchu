package server

import (
	"github.com/chuangxinyuan/fengchu/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewHTTPServer(
	logger log.Logger,
	config *config.Server,
) *http.Server {
	hc := config.Http
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			tracing.Server(),
			metadata.Server(),
		),
	}
	if hc.Network != "" {
		opts = append(opts, http.Network(hc.Network))
	}
	if hc.Addr != "" {
		opts = append(opts, http.Address(hc.Addr))
	}
	if hc.Timeout != nil {
		opts = append(opts, http.Timeout(hc.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	return srv
}
