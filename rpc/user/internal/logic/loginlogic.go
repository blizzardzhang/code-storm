package logic

import (
	"code-storm/rpc/user/internal/svc"
	"code-storm/rpc/user/userclient"
	"context"

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
	r, err := l.svcCtx.UserModel.FindByAccountAndPwd(l.ctx, in.Username, in.Password)
	if err != nil {
		return nil, err
	}

	return &userclient.LoginResp{
		UserName:         r.Account,
		Status:           "ok",
		AccessToken:      "",
		AccessExpire:     0,
		RefreshAfter:     0,
		CurrentAuthority: "admin",
		Id:               r.Id,
	}, nil
	/*var user usermodel.User
	err := l.svcCtx.DB.Take(&user, "account = ? and password = ?", in.Username, in.Password).Error
	if err != nil {
		return nil, errors.New("登录失败:" + err.Error())
	}
	return &userclient.LoginResp{
		UserName:         user.Account,
		Status:           "ok",
		AccessToken:      "",
		AccessExpire:     0,
		RefreshAfter:     0,
		CurrentAuthority: "admin",
		Id:               int64(user.ID),
	}, nil*/
}
