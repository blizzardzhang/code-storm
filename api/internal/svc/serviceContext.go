package svc

import (
	"code-storm/api/internal/config"
	"code-storm/api/internal/middleware"
	"code-storm/rpc/sys/client/appservice"
	"code-storm/rpc/sys/client/userservice"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	CheckUrl    rest.Middleware
	UserService userservice.UserService
	AppService  appservice.AppService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		CheckUrl:    middleware.NewCheckUrlMiddleware().Handle,
		UserService: userservice.NewUserService(zrpc.MustNewClient(c.SysRpc)),
		AppService:  appservice.NewAppService(zrpc.MustNewClient(c.SysRpc)),
	}
}
