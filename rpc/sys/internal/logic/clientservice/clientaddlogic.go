package clientservicelogic

import (
	"code-storm/rpc/model/sysmodel"
	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysclient"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClientAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientAddLogic {
	return &ClientAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClientAddLogic) ClientAdd(in *sysclient.AddClientReq) (*sysclient.AddClientResp, error) {
	/*client := sysmodel.StormClient{
		ClientId:       in.ClientId,
		ClientSecret:   in.ClientSecret,
		ClientKey:      in.ClientKey,
		Name:           in.Name,
		Remark:         in.Remark,
		AdditionalInfo: in.AdditionalInfo,
		RefreshAfter:   in.RefreshAfter,
		ExpireIn:       in.ExpireIn,
		GrantType:      in.GrantType,
	}
	_, err := l.svcCtx.ClientModel.Insert(l.ctx, &client)
	if err != nil {
		return nil, err
	}

	return &sysclient.AddClientResp{
		Data: "success",
	}, nil*/

	tx := l.svcCtx.DB.Save(&sysmodel.Client{
		ClientId:       in.ClientId,
		ClientSecret:   in.ClientSecret,
		ClientKey:      in.ClientKey,
		Name:           in.Name,
		Remark:         in.Remark,
		AdditionalInfo: in.AdditionalInfo,
		RefreshAfter:   in.RefreshAfter,
		ExpireIn:       in.ExpireIn,
		GrantType:      in.GrantType,
	})
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &sysclient.AddClientResp{
		Data: "insert success",
	}, nil

}
