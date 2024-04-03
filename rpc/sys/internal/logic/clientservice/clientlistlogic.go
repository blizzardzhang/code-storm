package clientservicelogic

import (
	"code-storm/rpc/model/sysmodel"
	"context"

	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysclient"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ClientListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClientListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientListLogic {
	return &ClientListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClientListLogic) ClientList(in *sysclient.ListClientReq) (*sysclient.ListClientResp, error) {
	current := in.Current
	size := in.PageSize
	var total int64
	var clients []sysmodel.Client
	tx := l.svcCtx.DB.Offset(int((current - 1) * size)).Limit(int(size)).Find(&clients)
	if tx.Error != nil {
		return nil, nil
	}
	//total
	tx = l.svcCtx.DB.Model(&sysmodel.Client{}).Count(&total)
	//data
	var data []*sysclient.ClientInfoResp
	for _, item := range clients {
		var client sysclient.ClientInfoResp
		_ = copier.Copy(&client, item)
		data = append(data, &client)
	}

	return &sysclient.ListClientResp{
		Total:     total,
		List:      data,
		PageSize:  size,
		Current:   current,
		TotalPage: total / size,
	}, nil
}
