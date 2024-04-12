package department

import (
	"code-storm/rpc/sys/client/departmentrpc"
	"context"
	"github.com/jinzhu/copier"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListDepartmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDepartmentLogic {
	return &ListDepartmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListDepartmentLogic) ListDepartment(req *types.ListDepartmentReq) (resp *types.ListDepartmentResp, err error) {
	res, err := l.svcCtx.DepartmentRpc.DepartmentList(l.ctx, &departmentrpc.ListDepartmentReq{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	var list []types.DepartmentInfoResp
	for _, item := range res.List {
		var department types.DepartmentInfoResp
		_ = copier.Copy(&department, item)
		list = append(list, department)
	}

	return &types.ListDepartmentResp{
		Records: list,
	}, nil
}
