package permissionservicelogic

import (
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPermissionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionDeleteLogic {
	return &PermissionDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PermissionDeleteLogic) PermissionDelete(in *sysClient.DeletePermissionReq) (*sysClient.DeletePermissionResp, error) {
	// todo: add your logic here and delete this line

	return &sysClient.DeletePermissionResp{}, nil
}
