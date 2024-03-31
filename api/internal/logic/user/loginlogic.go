package user

import (
	"code-storm/rpc/user/userclient"
	"context"

	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResp, err error) {
	// todo 登录逻辑
	res, err := l.svcCtx.UserService.Login(l.ctx, &userclient.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Status:           res.Status,
		CurrentAuthority: res.CurrentAuthority,
		Id:               res.Id,
		UserName:         res.UserName,
		AccessToken:      res.AccessToken,
		AccessExpire:     res.AccessExpire,
		RefreshAfter:     res.RefreshAfter,
	}, nil
}
