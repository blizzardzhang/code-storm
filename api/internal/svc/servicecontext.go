package svc

import (
	"code-storm/api/internal/config"
	"code-storm/rpc/sys/client/clientservice"
	"code-storm/rpc/sys/client/userservice"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	CheckUrl      rest.Middleware
	UserService   userservice.UserService
	ClientService clientservice.ClientService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserService:   userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		ClientService: clientservice.NewClientService(zrpc.MustNewClient(c.SysRpc)),
	}
}
