package svc

import (
	"code-storm/rpc/model/usermodel"
	"code-storm/rpc/user/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel usermodel.StormUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: usermodel.NewStormUserModel(sqlx.NewMysql(c.Mysql.Datasource), c.Cache),
	}
}
