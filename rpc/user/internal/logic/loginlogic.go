package logic

import (
	"context"

	"code-storm/rpc/user/internal/svc"
	"code-storm/rpc/user/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *userclient.LoginReq) (*userclient.LoginResp, error) {
	// todo: add your logic here and delete this line

	return &userclient.LoginResp{}, nil
}
