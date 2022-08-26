package boot

import (
	"github.com/chuangxinyuan/fengchu/component"
	"github.com/chuangxinyuan/fengchu/component/trace"
	"github.com/chuangxinyuan/fengchu/transport"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

var ProviderSet = wire.NewSet(
	component.ProviderSet,
	transport.ProviderSet,
)

type App struct {
	logger    *log.Helper
	trace     *trace.Tracer
	transport *transport.Transport
}

func NewApp(
	logger log.Logger,
	trace *trace.Tracer,
	transport *transport.Transport,
) *App {
	return &App{
		logger:    log.NewHelper(logger),
		trace:     trace,
		transport: transport,
	}
}

// Start 启动应用
func (a *App) Start() (err error) {
	// 设置 tracer
	if a.trace != nil {
		otel.SetTracerProvider(a.trace.TracerProvider())
		otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		))
	}
	err = a.transport.Start()
	return err
}

// Stop 停止应用
func (a *App) Stop() {
	// 关闭 transport 服务
	if err := a.transport.Stop(); err != nil {
		a.logger.Error(err)
	}
}

func (a *App) HttpServer() *http.Server {
	return a.transport.HttpServer()
}

func (a *App) GrpcServer() *grpc.Server {
	return a.transport.GrpcServer()
}
