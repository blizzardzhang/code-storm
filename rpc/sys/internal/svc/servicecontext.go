package svc

import (
	"code-storm/rpc/sys/internal/config"
	"code-storm/rpc/user/ent"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config config.Config
	DB     *ent.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)
	return &ServiceContext{
		Config: c,
		//ClientModel: sysmodel.NewStormClientModel(mysql, c.Cache),
		DB: db,
	}
}
