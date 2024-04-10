package permissionservicelogic

import (
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPermissionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionInfoLogic {
	return &PermissionInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PermissionInfoLogic) PermissionInfo(in *sysClient.PermissionInfoReq) (*sysClient.PermissionInfoResp, error) {
	// todo: add your logic here and delete this line

	return &sysClient.PermissionInfoResp{}, nil
}
