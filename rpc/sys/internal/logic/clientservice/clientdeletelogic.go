package clientservicelogic

import (
	"code-storm/rpc/model/sysmodel"
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClientDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientDeleteLogic {
	return &ClientDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClientDeleteLogic) ClientDelete(in *sysclient.DeleteClientReq) (*sysclient.DeleteClientResp, error) {
	tx := l.svcCtx.DB.Delete(&sysmodel.Client{}, "id in ?", in.Ids)
	if tx.Error != nil {
		return nil, nil
	}

	return &sysclient.DeleteClientResp{
		Data: "delete success",
	}, nil
}
