package permission

import (
	"code-storm/rpc/sys/client/permissionrpc"
	"context"
	"github.com/jinzhu/copier"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPermissionLogic {
	return &ListPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPermissionLogic) ListPermission(req *types.ListPermissionReq) (resp *types.ListPermissionResp, err error) {
	res, err := l.svcCtx.PermissionRpc.PermissionList(l.ctx, &permissionrpc.ListPermissionReq{
		Code:     req.Code,
		Category: req.Category,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	var list []types.PermissionDetailResp
	for _, item := range res.List {
		var perm types.PermissionDetailResp
		_ = copier.Copy(&perm, item)
		list = append(list, perm)
	}

	return &types.ListPermissionResp{
		Records: list,
	}, nil
}
