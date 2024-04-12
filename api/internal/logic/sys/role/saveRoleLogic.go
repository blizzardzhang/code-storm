package role

import (
	"code-storm/rpc/sys/client/rolerpc"
	"context"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveRoleLogic {
	return &SaveRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveRoleLogic) SaveRole(req *types.RoleAddReq) (resp *types.RoleAddResp, err error) {
	res, err := l.svcCtx.RoleRpc.RoleAdd(l.ctx, &rolerpc.AddRoleReq{
		Name:   req.Name,
		Code:   req.Code,
		Sort:   int64(req.Sort),
		Remark: req.Remark,
		Type:   int64(req.Type),
	})
	if err != nil {
		return nil, err
	}
	return &types.RoleAddResp{
		Data: res.Data,
	}, nil
}
