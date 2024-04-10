package roleservicelogic

import (
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleDeleteLogic) RoleDelete(in *sysClient.DeleteRoleReq) (*sysClient.DeleteRoleResp, error) {
	// todo: add your logic here and delete this line

	return &sysClient.DeleteRoleResp{}, nil
}
