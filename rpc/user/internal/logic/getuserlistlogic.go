package logic

import (
	"code-storm/rpc/user/internal/svc"
	"code-storm/rpc/user/userclient"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *userclient.UserListReq) (*userclient.UserListResp, error) {
	current := in.Current
	size := in.Size

	//page
	users, err := l.svcCtx.DB.User.Query().Where().Offset(int((current - 1) * size)).Limit(int(size)).All(l.ctx)
	if err != nil {
		return nil, err
	}
	// total
	total, err := l.svcCtx.DB.User.Query().Where().Count(l.ctx)
	if err != nil {
		return nil, err
	}

	resp := &userclient.UserListResp{}
	for _, item := range users {
		resp.List = append(resp.List, &userclient.UserInfo{
			Id:      item.ID.String(),
			Account: item.Account,
			Name:    item.Name,
			Phone:   item.Phone,
			Address: item.Address,
			//todo 时间转换
			//CreatedAt: item.CreateAt.UnixMilli(),
			RealName: item.RealName,
		})
	}

	resp.Total = uint64(total)

	return resp, nil
}
