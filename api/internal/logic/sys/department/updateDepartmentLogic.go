package department

import (
	"code-storm/rpc/sys/client/departmentrpc"
	"context"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDepartmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDepartmentLogic {
	return &UpdateDepartmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDepartmentLogic) UpdateDepartment(req *types.UpdateDepartmentReq) (resp *types.UpdateDepartmentResp, err error) {
	res, err := l.svcCtx.DepartmentRpc.DepartmentUpdate(l.ctx, &departmentrpc.UpdateDepartmentReq{
		Id:        req.Id,
		ParentId:  req.ParentId,
		Ancestors: req.Ancestors,
		Name:      req.Name,
		Sort:      int64(req.Sort),
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateDepartmentResp{
		Data: res.Data,
	}, nil
}
