package app

import (
	"code-storm/rpc/sys/client/apprpc"
	"context"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppDetailLogic {
	return &AppDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppDetailLogic) AppDetail(req *types.AppInfoReq) (resp *types.AppInfoResp, err error) {
	info, err := l.svcCtx.AppRpc.AppInfo(l.ctx, &apprpc.AppInfoReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.AppInfoResp{
		Id:                    info.Id,
		AppId:                 info.AppId,
		Name:                  info.Name,
		Key:                   info.Key,
		RedirectUri:           info.RedirectUri,
		Secret:                info.Secret,
		GrantType:             info.GrantType,
		AccessTokenValidity:   info.AccessTokenValidity,
		AdditionalInformation: info.AdditionalInfo,
		Status:                int(info.Status),
		CreateUser:            info.CreateUser,
		CreateAt:              info.CreateAt,
		UpdateUser:            info.UpdateUser,
	}, nil
}
