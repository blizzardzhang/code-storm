package svc

import (
	"code-storm/api/internal/config"
	"code-storm/api/internal/middleware"
	"code-storm/rpc/sys/client/apprpc"
	"code-storm/rpc/sys/client/departmentrpc"
	"code-storm/rpc/sys/client/permissionrpc"
	"code-storm/rpc/sys/client/rolerpc"
	"code-storm/rpc/sys/client/userrpc"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	CheckUrl      rest.Middleware
	UserRpc       userrpc.UserRpc
	AppRpc        apprpc.AppRpc
	RoleRpc       rolerpc.RoleRpc
	DepartmentRpc departmentrpc.DepartmentRpc
	PermissionRpc permissionrpc.PermissionRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		CheckUrl:      middleware.NewCheckUrlMiddleware().Handle,
		UserRpc:       userrpc.NewUserRpc(zrpc.MustNewClient(c.SysRpc)),
		AppRpc:        apprpc.NewAppRpc(zrpc.MustNewClient(c.SysRpc)),
		RoleRpc:       rolerpc.NewRoleRpc(zrpc.MustNewClient(c.SysRpc)),
		DepartmentRpc: departmentrpc.NewDepartmentRpc(zrpc.MustNewClient(c.SysRpc)),
		PermissionRpc: permissionrpc.NewPermissionRpc(zrpc.MustNewClient(c.SysRpc)),
	}
}
