package svc

import (
	"code-storm/common/enter"
	"code-storm/rpc/sys/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := enter.InitGorm(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		Db:     db,
	}

}
