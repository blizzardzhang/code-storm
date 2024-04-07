package config

import (
	"code-storm/common/enter"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Cache        cache.CacheConf
	DatabaseConf enter.DatabaseConf

	//ClientService clientservice.ClientService
	SysRpc zrpc.RpcClientConf
}
