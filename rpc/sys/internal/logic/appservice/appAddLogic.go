package appservicelogic

import (
	"code-storm/rpc/model/sys"
	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysClient"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type AppAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppAddLogic {
	return &AppAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppAddLogic) AppAdd(in *sysClient.AddAppReq) (*sysClient.AddAppResp, error) {
	//userId := l.ctx.Value("userId")
	app := sys.App{
		Name:                  in.Name,
		AppId:                 in.AppId,
		Key:                   in.Key,
		Secret:                in.Secret,
		GrantType:             in.GrantType,
		RedirectUri:           in.RedirectUri,
		AdditionalInformation: in.AdditionalInfo,
		AccessTokenValidity:   in.AccessTokenValidity,
		RefreshTokenValidity:  in.RefreshTokenValidity,
	}
	result := l.svcCtx.Db.Create(&app) //指针数据
	if result.Error != nil {
		err := errors.New("添加app失败:" + result.Error.Error())
		return nil, err
	}
	affected := result.RowsAffected

	return &sysClient.AddAppResp{
		Data: strconv.FormatInt(affected, 10),
	}, nil
}
