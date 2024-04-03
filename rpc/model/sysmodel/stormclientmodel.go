package sysmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StormClientModel = (*customStormClientModel)(nil)

type (
	// StormClientModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStormClientModel.
	StormClientModel interface {
		stormClientModel
	}

	customStormClientModel struct {
		*defaultStormClientModel
	}
)

// NewStormClientModel returns a model for the database table.
func NewStormClientModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) StormClientModel {
	return &customStormClientModel{
		defaultStormClientModel: newStormClientModel(conn, c, opts...),
	}
}
