package main

import (
	"flag"
	"fmt"

	"code-storm/rpc/sys/internal/config"
	apprpcServer "code-storm/rpc/sys/internal/server/apprpc"
	departmentrpcServer "code-storm/rpc/sys/internal/server/departmentrpc"
	permissionrpcServer "code-storm/rpc/sys/internal/server/permissionrpc"
	rolerpcServer "code-storm/rpc/sys/internal/server/rolerpc"
	userrpcServer "code-storm/rpc/sys/internal/server/userrpc"
	"code-storm/rpc/sys/internal/svc"
	"code-storm/rpc/sys/sysClient"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "rpc/sys/etc/sys.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		sysClient.RegisterUserRpcServer(grpcServer, userrpcServer.NewUserRpcServer(ctx))
		sysClient.RegisterAppRpcServer(grpcServer, apprpcServer.NewAppRpcServer(ctx))
		sysClient.RegisterDepartmentRpcServer(grpcServer, departmentrpcServer.NewDepartmentRpcServer(ctx))
		sysClient.RegisterRoleRpcServer(grpcServer, rolerpcServer.NewRoleRpcServer(ctx))
		sysClient.RegisterPermissionRpcServer(grpcServer, permissionrpcServer.NewPermissionRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
