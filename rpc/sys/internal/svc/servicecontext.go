package svc

import (
	"code-storm/common/enter"
	"code-storm/rpc/sys/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	//ClientModel sysmodel.StormClientModel
	DB *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	//mysql := sqlx.NewMysql(c.Mysql.Datasource)
	db := enter.InitGorm(c.Mysql.Datasource)
	return &ServiceContext{
		Config: c,
		//ClientModel: sysmodel.NewStormClientModel(mysql, c.Cache),
		DB: db,
	}
}
