package svc

import (
	"code-storm/rpc/sys/client/clientservice"
	"code-storm/rpc/user/ent"
	"code-storm/rpc/user/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	DB     *ent.Client

	ClientService clientservice.ClientService
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)
	return &ServiceContext{
		Config:        c,
		DB:            db,
		ClientService: clientservice.NewClientService(zrpc.MustNewClient(c.SysRpc)),
	}
}
