package departmentservicelogic

import (
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DepartmentInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDepartmentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DepartmentInfoLogic {
	return &DepartmentInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DepartmentInfoLogic) DepartmentInfo(in *sysClient.DepartmentInfoReq) (*sysClient.DepartmentInfoResp, error) {
	// todo: add your logic here and delete this line

	return &sysClient.DepartmentInfoResp{}, nil
}
