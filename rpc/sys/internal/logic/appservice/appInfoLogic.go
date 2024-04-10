package appservicelogic

import (
	"code-storm/rpc/model/sys"
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppInfoLogic {
	return &AppInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppInfoLogic) AppInfo(in *sysClient.AppInfoReq) (*sysClient.AppInfoResp, error) {
	var app sys.App
	result := l.svcCtx.Db.Find(&app, in.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &sysClient.AppInfoResp{
		Id:                   app.Id,
		AppId:                app.AppId,
		Name:                 app.Name,
		Secret:               app.Secret,
		GrantType:            app.GrantType,
		RedirectUri:          app.RedirectUri,
		AdditionalInfo:       app.AdditionalInformation,
		AccessTokenValidity:  app.AccessTokenValidity,
		RefreshTokenValidity: app.RefreshTokenValidity,
		CreateTime:           app.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateTime:           app.UpdateAt.Format("2006-01-02 15:04:05"),
	}, nil
}
