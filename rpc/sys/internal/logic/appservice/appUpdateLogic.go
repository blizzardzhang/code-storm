package appservicelogic

import (
	"code-storm/rpc/model/sys"
	"context"
	"strconv"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppUpdateLogic {
	return &AppUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppUpdateLogic) AppUpdate(in *sysClient.UpdateAppReq) (*sysClient.UpdateAppResp, error) {
	updates := map[string]interface{}{
		"name":                   in.Name,
		"app_id":                 in.AppId,
		"key":                    in.Key,
		"secret":                 in.Secret,
		"grant_type":             in.GrantType,
		"redirect_uri":           in.RedirectUri,
		"access_token_validity":  in.AccessTokenValidity,
		"refresh_token_validity": in.RefreshTokenValidity,
		"additional_information": in.AdditionalInfo,
	}
	var app sys.App
	result := l.svcCtx.Db.Model(&app).Where("id = ?", in.Id).Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}

	return &sysClient.UpdateAppResp{
		Data: strconv.FormatInt(result.RowsAffected, 10),
	}, nil
}
