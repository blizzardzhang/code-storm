package departmentservicelogic

import (
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DepartmentAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDepartmentAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DepartmentAddLogic {
	return &DepartmentAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DepartmentAddLogic) DepartmentAdd(in *sysClient.AddDepartmentReq) (*sysClient.AddDepartmentResp, error) {
	// todo: add your logic here and delete this line

	return &sysClient.AddDepartmentResp{}, nil
}
