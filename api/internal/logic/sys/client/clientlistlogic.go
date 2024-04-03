package client

import (
	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"
	"code-storm/rpc/sys/sysclient"
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ClientListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClientListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientListLogic {
	return &ClientListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClientListLogic) ClientList(req *types.ListClientReq) (resp *types.ListClientResp, err error) {
	list, err := l.svcCtx.ClientService.ClientList(l.ctx, &sysclient.ListClientReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	var data []types.ClientData
	for _, item := range list.List {
		var d types.ClientData
		_ = copier.Copy(&d, item)
		data = append(data, d)
	}

	return &types.ListClientResp{
		Total:     list.Total,
		Data:      data,
		TotalPage: list.TotalPage,
		Current:   list.Current,
		PageSize:  list.PageSize,
	}, nil
}
