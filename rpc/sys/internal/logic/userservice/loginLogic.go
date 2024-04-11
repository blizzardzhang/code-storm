package userservicelogic

import (
	"code-storm/common/tool"
	"code-storm/rpc/model/sys"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"time"

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
	//1.用户查询
	var user sys.User
	err := l.svcCtx.Db.Take(&user, "account = ?", in.Account).Error
	if err != nil {
		err = errors.New("用户名或密码错误")
		return nil, errors.Wrap(err, "LoginLogic.Login")
	}

	//2.校验密码
	if !tool.CheckPwd(user.Password, in.Password) {
		err = errors.New("用户名或密码错误")
		return nil, errors.Wrap(err, "LoginLogic.Login")
	}
	//3.查询app 根据app获取token有效期
	var app sys.App
	tx := l.svcCtx.Db.Take(&app, "app_id = ?", in.AppId)
	if tx.Error != nil {
		err = errors.New("app不存在")
		return nil, errors.Wrap(err, "LoginLogic.Login")
	}

	now := time.Now().Unix()
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret
	validity := app.AccessTokenValidity
	userId := user.Id
	account := user.Account
	token, err := l.getJwtToken(accessSecret, now, validity, userId, account)
	if err != nil {
		return nil, err
	}

	return &sysClient.LoginResp{
		Id:           user.Id,
		Account:      user.Account,
		AccessToken:  token,
		AccessExpire: validity,
		RefreshAfter: 0,
	}, nil
}

// 生成jwt的token
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds int64, userId string, account string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	claims["account"] = account
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
