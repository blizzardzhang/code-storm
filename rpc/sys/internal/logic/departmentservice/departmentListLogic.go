package departmentservicelogic

import (
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DepartmentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDepartmentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DepartmentListLogic {
	return &DepartmentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DepartmentListLogic) DepartmentList(in *sysClient.ListDepartmentReq) (*sysClient.ListDepartmentResp, error) {
	// todo: add your logic here and delete this line

	return &sysClient.ListDepartmentResp{}, nil
}
