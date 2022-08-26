package redis

import (
	"context"
	"github.com/chuangxinyuan/fengchu/config"
	"github.com/spf13/cast"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

// New 创建 redis 客户端
func New(config *config.Redis, logger log.Logger) (*redis.Client, func(), error) {
	if config == nil {
		return nil, func() {}, nil
	}
	option := &redis.Options{
		Addr: config.Address,
	}
	if config.Username != "" {
		option.Username = config.Username
	}
	if config.Password != "" {
		option.Password = config.Password
	}
	if config.Database != 0 {
		option.DB = cast.ToInt(config.Database)
	}
	if config.ReadTimeout.AsDuration() != 0 {
		option.ReadTimeout = config.ReadTimeout.AsDuration() * time.Second
	}
	if config.WriteTimeout.AsDuration() != 0 {
		option.WriteTimeout = config.WriteTimeout.AsDuration() * time.Second
	}

	client := redis.NewClient(option)

	ctx := context.Background()

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the redis client")

		if err := client.Close(); err != nil {
			log.NewHelper(logger).Error(err.Error())
		}
	}

	return client, cleanup, nil
}
