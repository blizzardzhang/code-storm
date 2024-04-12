package role

import (
	"code-storm/rpc/sys/client/rolerpc"
	"context"
	"github.com/jinzhu/copier"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRoleLogic {
	return &ListRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListRoleLogic) ListRole(req *types.RoleListReq) (resp *types.RoleListResp, err error) {
	res, err := l.svcCtx.RoleRpc.RoleList(l.ctx, &rolerpc.ListRoleReq{
		PageSize: int64(req.PageSize),
		Current:  int64(req.Current),
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	var list []types.RoleInfoResp
	for _, item := range res.List {
		var role types.RoleInfoResp
		_ = copier.Copy(&role, item)
		list = append(list, role)
	}
	return &types.RoleListResp{
		Records:   list,
		Current:   int(res.Current),
		PageSize:  int(res.PageSize),
		Total:     int(res.Total),
		TotalPage: int(res.TotalPage),
	}, nil
}
