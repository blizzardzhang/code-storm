package app

import (
	"code-storm/rpc/sys/client/apprpc"
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
	res, err := l.svcCtx.AppRpc.AppList(l.ctx, &apprpc.ListAppReq{
		Current:  int64(req.Current),
		PageSize: int64(req.PageSize),
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	var list []types.AppInfoResp
	for _, item := range res.List {
		var app types.AppInfoResp
		_ = copier.Copy(&app, item)
		list = append(list, app)
	}
	return &types.ListAppResp{
		Records:   list,
		Current:   int(res.Current),
		PageSize:  int(res.PageSize),
		Total:     int(res.Total),
		TotalPage: int(res.TotalPage),
	}, nil
}
