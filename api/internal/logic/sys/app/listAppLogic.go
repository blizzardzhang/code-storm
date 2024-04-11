package app

import (
	"code-storm/rpc/sys/sysClient"
	"context"
	"github.com/jinzhu/copier"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAppLogic {
	return &ListAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAppLogic) ListApp(req *types.ListAppReq) (resp *types.ListAppResp, err error) {
	res, err := l.svcCtx.AppService.AppList(l.ctx, &sysClient.ListAppReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	var list []types.AppData
	for _, item := range res.List {
		var app types.AppData
		_ = copier.Copy(&app, &item)
		list = append(list, app)
	}
	return &types.ListAppResp{
		Data:      list,
		Current:   res.Current,
		PageSize:  res.PageSize,
		Total:     res.Total,
		TotalPage: res.TotalPage,
	}, nil
}
