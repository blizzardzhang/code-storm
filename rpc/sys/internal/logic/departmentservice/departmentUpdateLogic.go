package departmentservicelogic

import (
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DepartmentUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDepartmentUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DepartmentUpdateLogic {
	return &DepartmentUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DepartmentUpdateLogic) DepartmentUpdate(in *sysClient.UpdateDepartmentReq) (*sysClient.UpdateDepartmentResp, error) {
	// todo: add your logic here and delete this line

	return &sysClient.UpdateDepartmentResp{}, nil
}
