package role

import (
	"code-storm/rpc/sys/client/rolerpc"
	"context"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRoleLogic) DeleteRole(req *types.RoleDeleteReq) (resp *types.RoleDeleteResp, err error) {
	res, err := l.svcCtx.RoleRpc.RoleDelete(l.ctx, &rolerpc.DeleteRoleReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.RoleDeleteResp{
		Data: res.Data,
	}, nil
}
