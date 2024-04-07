package clientservicelogic

import (
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
	return nil, nil
}
