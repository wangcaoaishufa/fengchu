package component

import (
	"github.com/chuangxinyuan/fengchu/component/config"
	"github.com/chuangxinyuan/fengchu/component/log"
	"github.com/chuangxinyuan/fengchu/component/orm"
	"github.com/chuangxinyuan/fengchu/component/redis"
	"github.com/chuangxinyuan/fengchu/component/trace"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	orm.New,
	redis.New,
	trace.New,
)
