package clientservicelogic

import (
	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysclient"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClientAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientAddLogic {
	return &ClientAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClientAddLogic) ClientAdd(in *sysclient.AddClientReq) (*sysclient.AddClientResp, error) {
	return nil, nil
}
