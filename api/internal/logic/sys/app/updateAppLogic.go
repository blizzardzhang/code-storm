package app

import (
	"code-storm/rpc/sys/client/apprpc"
	"context"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAppLogic {
	return &UpdateAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAppLogic) UpdateApp(req *types.UpdateAppReq) (resp *types.UpdateAppResp, err error) {
	res, err := l.svcCtx.AppRpc.AppUpdate(l.ctx, &apprpc.UpdateAppReq{
		Id:                  req.Id,
		Name:                req.Name,
		GrantType:           req.GrantType,
		AccessTokenValidity: req.AccessTokenValidity,
		AdditionalInfo:      req.AdditionalInformation,
		Key:                 req.Key,
		RedirectUri:         req.RedirectUri,
		AppId:               req.AppId,
		Secret:              req.Secret,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateAppResp{
		Data: res.Data,
	}, nil
}
