package client

import (
	"code-storm/rpc/sys/sysclient"
	"context"
	"github.com/pkg/errors"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClientAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientAddLogic {
	return &ClientAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClientAddLogic) ClientAdd(req *types.AddClientReq) (resp *types.AddClientResp, err error) {

	r, err := l.svcCtx.ClientService.ClientAdd(l.ctx, &sysclient.AddClientReq{
		ClientSecret:   req.ClientSecret,
		AdditionalInfo: req.AdditionalInfo,
		RefreshAfter:   req.RefreshAfter,
		Name:           req.Name,
		Remark:         req.Remark,
		ClientKey:      req.ClientKey,
		GrantType:      req.GrantType,
		ExpireIn:       req.ExpireIn,
		ClientId:       req.ClientId,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "ClientAdd err:%v", err)
	}

	return &types.AddClientResp{
		Data: r.Data,
	}, nil
}
