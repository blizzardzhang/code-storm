package departmentservicelogic

import (
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DepartmentDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDepartmentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DepartmentDeleteLogic {
	return &DepartmentDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DepartmentDeleteLogic) DepartmentDelete(in *sysClient.DeleteDepartmentReq) (*sysClient.DeleteDepartmentResp, error) {
	// todo: add your logic here and delete this line

	return &sysClient.DeleteDepartmentResp{}, nil
}
