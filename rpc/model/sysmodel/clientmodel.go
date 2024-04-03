package sysmodel

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Client struct {
	Id             string                `gorm:"primaryKey" json:"id"`
	Name           string                `json:"name"`
	ClientId       string                `json:"clientId"`
	ClientSecret   string                `json:"clientSecret"`
	ClientKey      string                `json:"clientKey"`
	AdditionalInfo string                `json:"additionalInfo"`
	GrantType      string                `json:"grantType"`
	ExpireIn       int64                 `json:"expireIn"`
	RefreshAfter   int64                 `json:"refreshAfter"`
	Remark         string                `json:"remark"`
	Status         int64                 `json:"status"`
	DeletedAt      soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name"`
	CreatedAt      time.Time             `gorm:"autoCreateTime"`
	UpdatedAt      time.Time             `gorm:"autoUpdateTime"`
}
