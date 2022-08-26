package log

import (
	"github.com/chuangxinyuan/fengchu/config"
	kzap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/google/wire"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ProviderSet = wire.NewSet(NewZapLogger, NewKratosLogger)

func NewZapLogger(config *config.Zap) *zap.Logger {
	cores := Zap.GetZapCores(config)
	zapLogger := zap.New(zapcore.NewTee(cores...))
	return zapLogger
}

func NewKratosLogger(config *config.App, zLogger *zap.Logger) log.Logger {
	logger := log.With(
		kzap.NewLogger(zLogger),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", config.AppId,
		"service.name", config.AppName,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
	return logger
}
