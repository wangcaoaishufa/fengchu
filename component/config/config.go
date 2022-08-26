package config

import (
	systemConfig "github.com/chuangxinyuan/fengchu/config"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewConfig,
	NewSystemConfig,
	wire.FieldsOf(
		new(*systemConfig.SystemConfig),
		"App",
		"Database",
		"Redis",
		"Zap",
		"Trace",
	),
)

func NewConfig(paths []string) config.Config {
	var sources []config.Source
	if len(paths) > 0 {
		for _, path := range paths {
			sources = append(sources, file.NewSource(path))
		}
	}
	c := config.New(
		config.WithSource(sources...),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}
	return c
}

func NewSystemConfig(c config.Config) (systemConfig *systemConfig.SystemConfig) {
	if err := c.Scan(&systemConfig); err != nil {
		panic(err)
	}
	return systemConfig
}
