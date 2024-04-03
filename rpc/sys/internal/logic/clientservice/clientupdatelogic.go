package clientservicelogic

import (
	"code-storm/rpc/model/sysmodel"
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClientUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientUpdateLogic {
	return &ClientUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClientUpdateLogic) ClientUpdate(in *sysclient.UpdateClientReq) (*sysclient.UpdateClientResp, error) {
	tx := l.svcCtx.DB.Save(&sysmodel.Client{
		Id:     in.Id,
		Name:   in.Name,
		Remark: in.Remark,
		Status: in.Status,
	})
	if tx.Error != nil {
		return nil, nil
	}

	return &sysclient.UpdateClientResp{
		Data: "update success",
	}, nil
}
