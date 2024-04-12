package app

import (
	"code-storm/rpc/sys/client/apprpc"
	"context"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveAppLogic {
	return &SaveAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveAppLogic) SaveApp(req *types.AddAppReq) (resp *types.AddAppResp, err error) {
	//userId := l.ctx.Value("userId").(string)
	res, err := l.svcCtx.AppRpc.AppAdd(l.ctx, &apprpc.AddAppReq{
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
