package config

import (
	"code-storm/common/enter"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	//Mysql struct {
	//	Datasource string
	//}
	Cache        cache.CacheConf
	DatabaseConf enter.DatabaseConf
}
