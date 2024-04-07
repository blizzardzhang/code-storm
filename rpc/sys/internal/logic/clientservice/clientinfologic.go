package clientservicelogic

import (
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClientInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientInfoLogic {
	return &ClientInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClientInfoLogic) ClientInfo(in *sysclient.ClientInfoReq) (*sysclient.ClientInfoResp, error) {
	return nil, nil
}
