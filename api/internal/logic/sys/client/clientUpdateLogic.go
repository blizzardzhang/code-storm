package client

import (
	"context"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClientUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientUpdateLogic {
	return &ClientUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClientUpdateLogic) ClientUpdate(req *types.UpdateClientReq) (resp *types.UpdateClientResp, err error) {
	// todo: add your logic here and delete this line

	return
}
