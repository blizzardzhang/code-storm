package client

import (
	"code-storm/rpc/sys/sysclient"
	"context"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClientInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientInfoLogic {
	return &ClientInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClientInfoLogic) ClientInfo(req *types.ClientInfoReq) (resp *types.ClientData, err error) {
	info, err := l.svcCtx.ClientService.ClientInfo(l.ctx, &sysclient.ClientInfoReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.ClientData{
		Id:             info.Id,
		ClientId:       info.ClientId,
		GrantType:      info.GrantType,
		ExpireIn:       info.ExpireIn,
		Name:           info.Name,
		RefreshAfter:   info.RefreshAfter,
		Status:         info.Status,
		ClientKey:      info.ClientKey,
		IsDeleted:      info.IsDeleted,
		Remark:         info.Remark,
		AdditionalInfo: info.AdditionalInfo,
	}, nil
}
