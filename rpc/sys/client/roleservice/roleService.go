// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package roleservice

import (
	"context"

	"code-storm/rpc/sys/sysClient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddAppReq            = sysClient.AddAppReq
	AddAppResp           = sysClient.AddAppResp
	AddDepartmentReq     = sysClient.AddDepartmentReq
	AddDepartmentResp    = sysClient.AddDepartmentResp
	AddPermissionReq     = sysClient.AddPermissionReq
	AddPermissionResp    = sysClient.AddPermissionResp
	AddRoleReq           = sysClient.AddRoleReq
	AddRoleResp          = sysClient.AddRoleResp
	AppInfoReq           = sysClient.AppInfoReq
	AppInfoResp          = sysClient.AppInfoResp
	DeleteAppReq         = sysClient.DeleteAppReq
	DeleteAppResp        = sysClient.DeleteAppResp
	DeleteDepartmentReq  = sysClient.DeleteDepartmentReq
	DeleteDepartmentResp = sysClient.DeleteDepartmentResp
	DeletePermissionReq  = sysClient.DeletePermissionReq
	DeletePermissionResp = sysClient.DeletePermissionResp
	DeleteRoleReq        = sysClient.DeleteRoleReq
	DeleteRoleResp       = sysClient.DeleteRoleResp
	DepartmentInfoReq    = sysClient.DepartmentInfoReq
	DepartmentInfoResp   = sysClient.DepartmentInfoResp
	ListAppReq           = sysClient.ListAppReq
	ListAppResp          = sysClient.ListAppResp
	ListDepartmentReq    = sysClient.ListDepartmentReq
	ListDepartmentResp   = sysClient.ListDepartmentResp
	ListPermissionReq    = sysClient.ListPermissionReq
	ListPermissionResp   = sysClient.ListPermissionResp
	ListRoleReq          = sysClient.ListRoleReq
	ListRoleResp         = sysClient.ListRoleResp
	LoginReq             = sysClient.LoginReq
	LoginResp            = sysClient.LoginResp
	PermissionInfoReq    = sysClient.PermissionInfoReq
	PermissionInfoResp   = sysClient.PermissionInfoResp
	RoleInfoReq          = sysClient.RoleInfoReq
	RoleInfoResp         = sysClient.RoleInfoResp
	UpdateAppReq         = sysClient.UpdateAppReq
	UpdateAppResp        = sysClient.UpdateAppResp
	UpdateDepartmentReq  = sysClient.UpdateDepartmentReq
	UpdateDepartmentResp = sysClient.UpdateDepartmentResp
	UpdatePermissionReq  = sysClient.UpdatePermissionReq
	UpdatePermissionResp = sysClient.UpdatePermissionResp
	UpdateRoleReq        = sysClient.UpdateRoleReq
	UpdateRoleResp       = sysClient.UpdateRoleResp
	UserInfo             = sysClient.UserInfo
	UserListReq          = sysClient.UserListReq
	UserListResp         = sysClient.UserListResp

	RoleService interface {
		RoleAdd(ctx context.Context, in *AddRoleReq, opts ...grpc.CallOption) (*AddRoleResp, error)
		RoleInfo(ctx context.Context, in *RoleInfoReq, opts ...grpc.CallOption) (*RoleInfoResp, error)
		RoleList(ctx context.Context, in *ListRoleReq, opts ...grpc.CallOption) (*ListRoleResp, error)
		RoleUpdate(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleResp, error)
		RoleDelete(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*DeleteRoleResp, error)
	}

	defaultRoleService struct {
		cli zrpc.Client
	}
)

func NewRoleService(cli zrpc.Client) RoleService {
	return &defaultRoleService{
		cli: cli,
	}
}

func (m *defaultRoleService) RoleAdd(ctx context.Context, in *AddRoleReq, opts ...grpc.CallOption) (*AddRoleResp, error) {
	client := sysClient.NewRoleServiceClient(m.cli.Conn())
	return client.RoleAdd(ctx, in, opts...)
}

func (m *defaultRoleService) RoleInfo(ctx context.Context, in *RoleInfoReq, opts ...grpc.CallOption) (*RoleInfoResp, error) {
	client := sysClient.NewRoleServiceClient(m.cli.Conn())
	return client.RoleInfo(ctx, in, opts...)
}

func (m *defaultRoleService) RoleList(ctx context.Context, in *ListRoleReq, opts ...grpc.CallOption) (*ListRoleResp, error) {
	client := sysClient.NewRoleServiceClient(m.cli.Conn())
	return client.RoleList(ctx, in, opts...)
}

func (m *defaultRoleService) RoleUpdate(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleResp, error) {
	client := sysClient.NewRoleServiceClient(m.cli.Conn())
	return client.RoleUpdate(ctx, in, opts...)
}

func (m *defaultRoleService) RoleDelete(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*DeleteRoleResp, error) {
	client := sysClient.NewRoleServiceClient(m.cli.Conn())
	return client.RoleDelete(ctx, in, opts...)
}
