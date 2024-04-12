package app

import (
	"code-storm/rpc/sys/sysClient"
	"context"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppInfoLogic {
	return &AppInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppInfoLogic) AppInfo(req *types.AppInfoReq) (resp *types.AppData, err error) {
	info, err := l.svcCtx.AppService.AppInfo(l.ctx, &sysClient.AppInfoReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.AppData{
		Id:                    info.Id,
		AppId:                 info.AppId,
		Name:                  info.Name,
		Key:                   info.Key,
		RedirectUri:           info.RedirectUri,
		Secret:                info.Secret,
		GrantType:             info.GrantType,
		AccessTokenValidity:   info.AccessTokenValidity,
		AdditionalInformation: info.AdditionalInfo,
		Status:                info.Status,
		CreateUser:            info.CreateUser,
		CreateAt:              info.CreateAt,
		UpdateUser:            info.UpdateUser,
	}, nil
}
