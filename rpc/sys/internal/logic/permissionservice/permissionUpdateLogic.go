package permissionservicelogic

import (
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPermissionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionUpdateLogic {
	return &PermissionUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PermissionUpdateLogic) PermissionUpdate(in *sysClient.UpdatePermissionReq) (*sysClient.UpdatePermissionResp, error) {
	// todo: add your logic here and delete this line

	return &sysClient.UpdatePermissionResp{}, nil
}
