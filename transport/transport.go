package transport

import (
	"github.com/chuangxinyuan/fengchu/config"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

var hostname, _ = os.Hostname()

var ProviderSet = wire.NewSet(
	//server.NewGRPCServer,
	//server.NewHTTPServer,
	New,
)

type Transport struct {
	logger     *log.Helper
	server     *kratos.App
	httpServer *http.Server
	grpcServer *grpc.Server
}

func New(
	logger log.Logger,
	config *config.App,
	httpServer *http.Server,
	grpcServer *grpc.Server,
) *Transport {
	var servers []transport.Server
	if httpServer != nil {
		servers = append(servers, httpServer)
	}
	if grpcServer != nil {
		servers = append(servers, grpcServer)
	}

	options := []kratos.Option{
		kratos.ID(hostname),
		kratos.Name(config.AppName),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(servers...),
	}

	server := kratos.New(options...)

	return &Transport{
		logger:     log.NewHelper(logger),
		server:     server,
		httpServer: httpServer,
		grpcServer: grpcServer,
	}
}

func (t *Transport) Start() error {
	t.logger.Info("transport server starting ...")
	if err := t.server.Run(); err != nil {
		return err
	}
	return nil
}

func (t *Transport) Stop() error {
	if err := t.server.Stop(); err != nil {
		return err
	}

	t.logger.Info("transport server stopping ...")
	return nil
}

func (t *Transport) HttpServer() *http.Server {
	return t.httpServer
}

func (t *Transport) GrpcServer() *grpc.Server {
	return t.grpcServer
}
