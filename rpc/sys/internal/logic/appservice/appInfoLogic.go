package appservicelogic

import (
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppInfoLogic {
	return &AppInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppInfoLogic) AppInfo(in *sysClient.AppInfoReq) (*sysClient.AppInfoResp, error) {
	// todo: add your logic here and delete this line

	return &sysClient.AppInfoResp{}, nil
}
