package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/ent"
	_ "github.com/saas-mingyang/mingyang-admin-core/rpc/ent/runtime"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config config.Config
	DB     *ent.Client
	Redis  redis.UniversalClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	entOpts := []ent.Option{
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
	}
	if c.DatabaseConf.Debug {
		entOpts = append(entOpts,
			ent.Debug(),
			ent.Log(func(a ...interface{}) {
				logx.Info(a...)
			}),
		)
	}
	db := ent.NewClient(entOpts...)
	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  c.RedisConf.MustNewUniversalRedis(),
	}
}
