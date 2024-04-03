package svc

import (
	"code-storm/rpc/model/usermodel"
	"code-storm/rpc/user/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel usermodel.StormUserModel
	//DB *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := sqlx.NewMysql(c.Mysql.Datasource)
	//db := enter.InitGorm(c.Mysql.Datasource)
	return &ServiceContext{
		Config:    c,
		UserModel: usermodel.NewStormUserModel(mysql, c.Cache),
		//DB: db,
	}
}
