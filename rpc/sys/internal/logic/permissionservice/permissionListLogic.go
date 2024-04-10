package permissionservicelogic

import (
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPermissionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionListLogic {
	return &PermissionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PermissionListLogic) PermissionList(in *sysClient.ListPermissionReq) (*sysClient.ListPermissionResp, error) {
	// todo: add your logic here and delete this line

	return &sysClient.ListPermissionResp{}, nil
}
