package app

import (
	"code-storm/rpc/sys/sysClient"
	"context"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAppLogic {
	return &DeleteAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAppLogic) DeleteApp(req *types.DeleteAppReq) (resp *types.DeleteAppResp, err error) {
	res, err := l.svcCtx.AppService.AppDelete(l.ctx, &sysClient.DeleteAppReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}
	return &types.DeleteAppResp{
		Data: res.Data,
	}, nil
}
