package client

import (
	"code-storm/rpc/sys/sysclient"
	"context"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClientDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientDeleteLogic {
	return &ClientDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClientDeleteLogic) ClientDelete(req *types.DeleteClientReq) (resp *types.DeleteClientResp, err error) {
	r, err := l.svcCtx.ClientService.ClientDelete(l.ctx, &sysclient.DeleteClientReq{
		Ids: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.DeleteClientResp{
		Data: r.Data,
	}, nil
}
