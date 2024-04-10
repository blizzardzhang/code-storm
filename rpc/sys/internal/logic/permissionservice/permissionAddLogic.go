package permissionservicelogic

import (
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPermissionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionAddLogic {
	return &PermissionAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PermissionAddLogic) PermissionAdd(in *sysClient.AddPermissionReq) (*sysClient.AddPermissionResp, error) {
	// todo: add your logic here and delete this line

	return &sysClient.AddPermissionResp{}, nil
}
