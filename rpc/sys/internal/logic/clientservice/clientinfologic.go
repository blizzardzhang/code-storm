package clientservicelogic

import (
	"code-storm/rpc/model/sysmodel"
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClientInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientInfoLogic {
	return &ClientInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClientInfoLogic) ClientInfo(in *sysclient.ClientInfoReq) (*sysclient.ClientInfoResp, error) {
	var data sysmodel.Client
	first := l.svcCtx.DB.First(&data, in.Id)
	if first.Error != nil {
		return nil, nil
	}
	return &sysclient.ClientInfoResp{
		Id:             data.Id,
		Name:           data.Name,
		AdditionalInfo: data.AdditionalInfo,
		RefreshAfter:   data.RefreshAfter,
		Remark:         data.Remark,
		IsDeleted:      int64(data.DeletedAt),
		ClientId:       data.ClientId,
		ExpireIn:       data.ExpireIn,
		ClientKey:      data.ClientKey,
		ClientSecret:   data.ClientSecret,
		GrantType:      data.GrantType,
	}, nil
}
