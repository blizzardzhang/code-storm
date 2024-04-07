package clientservicelogic

import (
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClientUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientUpdateLogic {
	return &ClientUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClientUpdateLogic) ClientUpdate(in *sysclient.UpdateClientReq) (*sysclient.UpdateClientResp, error) {
	return nil, nil
}
