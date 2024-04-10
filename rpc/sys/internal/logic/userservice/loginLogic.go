package userservicelogic

import (
	"code-storm/common/tool"
	"code-storm/rpc/model/sys"
	"context"
	"github.com/pkg/errors"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysClient"

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

func (l *LoginLogic) Login(in *sysClient.LoginReq) (*sysClient.LoginResp, error) {
	var user sys.User
	err := l.svcCtx.Db.Take(&user, "account = ?", in.Account).Error
	if err != nil {
		err = errors.New("用户名或密码错误")
		return nil, errors.Wrap(err, "LoginLogic.Login")
	}

	//校验密码
	if !tool.CheckPwd(user.Password, in.Password) {
		err = errors.New("用户名或密码错误")
		return nil, errors.Wrap(err, "LoginLogic.Login")
	}
	token, err := tool.GenToken(tool.JwtPayLoad{
		UserId:  user.Id,
		Account: user.Account,
		Role:    "admin",
	}, l.svcCtx.Config.JwtAuth.AccessSecret, l.svcCtx.Config.JwtAuth.AccessExpire)
	if err != nil {
		logx.Error(err)
		err = errors.New("jwt错误")
		return nil, err
	}

	return &sysClient.LoginResp{
		Id:           user.Id,
		Account:      user.Account,
		AccessToken:  token,
		AccessExpire: 6499,
		RefreshAfter: 0,
	}, nil
}
