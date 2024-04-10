package main

import (
	"flag"
	"fmt"

	"code-storm/rpc/sys/internal/config"
	appserviceServer "code-storm/rpc/sys/internal/server/appservice"
	departmentserviceServer "code-storm/rpc/sys/internal/server/departmentservice"
	permissionserviceServer "code-storm/rpc/sys/internal/server/permissionservice"
	roleserviceServer "code-storm/rpc/sys/internal/server/roleservice"
	userserviceServer "code-storm/rpc/sys/internal/server/userservice"
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
		sysClient.RegisterUserServiceServer(grpcServer, userserviceServer.NewUserServiceServer(ctx))
		sysClient.RegisterAppServiceServer(grpcServer, appserviceServer.NewAppServiceServer(ctx))
		sysClient.RegisterDepartmentServiceServer(grpcServer, departmentserviceServer.NewDepartmentServiceServer(ctx))
		sysClient.RegisterRoleServiceServer(grpcServer, roleserviceServer.NewRoleServiceServer(ctx))
		sysClient.RegisterPermissionServiceServer(grpcServer, permissionserviceServer.NewPermissionServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
