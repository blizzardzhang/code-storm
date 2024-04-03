package usermodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StormUserModel = (*customStormUserModel)(nil)

type (
	// StormUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStormUserModel.
	StormUserModel interface {
		stormUserModel
	}

	customStormUserModel struct {
		*defaultStormUserModel
	}
)

// NewStormUserModel returns a model for the database table.
func NewStormUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) StormUserModel {
	return &customStormUserModel{
		defaultStormUserModel: newStormUserModel(conn, c, opts...),
	}
}
