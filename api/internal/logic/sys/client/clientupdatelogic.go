package client

import (
	"code-storm/rpc/sys/sysclient"
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
	r, err := l.svcCtx.ClientService.ClientUpdate(l.ctx, &sysclient.UpdateClientReq{
		Id:     req.Id,
		Name:   req.Name,
		Remark: req.Remark,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateClientResp{
		Data: r.Data,
	}, nil
}
