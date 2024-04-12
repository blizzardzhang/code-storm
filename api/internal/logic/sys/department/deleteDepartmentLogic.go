package department

import (
	"code-storm/rpc/sys/client/departmentrpc"
	"context"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDepartmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDepartmentLogic {
	return &DeleteDepartmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDepartmentLogic) DeleteDepartment(req *types.DeleteDepartmentReq) (resp *types.DeleteDepartmentResp, err error) {
	res, err := l.svcCtx.DepartmentRpc.DepartmentDelete(l.ctx, &departmentrpc.DeleteDepartmentReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}
	return &types.DeleteDepartmentResp{
		Data: res.Data,
	}, nil
}
