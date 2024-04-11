package app

import (
	"code-storm/rpc/sys/client/appservice"
	"context"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAppLogic {
	return &AddAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAppLogic) AddApp(req *types.AddAppReq) (resp *types.AddAppResp, err error) {
	//userId := l.ctx.Value("userId").(string)
	res, err := l.svcCtx.AppService.AppAdd(l.ctx, &appservice.AddAppReq{
		AppId:                req.AppId,
		RedirectUri:          req.RedirectUri,
		AccessTokenValidity:  req.AccessTokenValidity,
		Key:                  req.Key,
		Secret:               req.Secret,
		RefreshTokenValidity: req.RefreshTokenValidity,
		Name:                 req.Name,
		GrantType:            req.GrantType,
		AdditionalInfo:       req.AdditionalInformation,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddAppResp{
		Data: res.Data,
	}, nil
}
