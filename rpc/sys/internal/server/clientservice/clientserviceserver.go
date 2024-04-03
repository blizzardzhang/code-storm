// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"code-storm/rpc/sys/internal/logic/clientservice"
	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysclient"
)

type ClientServiceServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedClientServiceServer
}

func NewClientServiceServer(svcCtx *svc.ServiceContext) *ClientServiceServer {
	return &ClientServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *ClientServiceServer) ClientAdd(ctx context.Context, in *sysclient.AddClientReq) (*sysclient.AddClientResp, error) {
	l := clientservicelogic.NewClientAddLogic(ctx, s.svcCtx)
	return l.ClientAdd(in)
}

func (s *ClientServiceServer) ClientInfo(ctx context.Context, in *sysclient.ClientInfoReq) (*sysclient.ClientInfoResp, error) {
	l := clientservicelogic.NewClientInfoLogic(ctx, s.svcCtx)
	return l.ClientInfo(in)
}

func (s *ClientServiceServer) ClientList(ctx context.Context, in *sysclient.ListClientReq) (*sysclient.ListClientResp, error) {
	l := clientservicelogic.NewClientListLogic(ctx, s.svcCtx)
	return l.ClientList(in)
}

func (s *ClientServiceServer) ClientUpdate(ctx context.Context, in *sysclient.UpdateClientReq) (*sysclient.UpdateClientResp, error) {
	l := clientservicelogic.NewClientUpdateLogic(ctx, s.svcCtx)
	return l.ClientUpdate(in)
}

func (s *ClientServiceServer) ClientDelete(ctx context.Context, in *sysclient.DeleteClientReq) (*sysclient.DeleteClientResp, error) {
	l := clientservicelogic.NewClientDeleteLogic(ctx, s.svcCtx)
	return l.ClientDelete(in)
}
