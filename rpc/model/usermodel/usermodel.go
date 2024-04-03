package usermodel

import (
	"time"
)

type User struct {
	Id        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
	Account   string    `gorm:"unique;size:32" json:"account"`
	Password  string    `gorm:"size:64" json:"password"`
	Name      string    `gorm:"size:32" json:"name"`
	RealName  string    `gorm:"size:32" json:"realName"`
	Phone     string    `gorm:"size:32" json:"phone"`
	UserType  int       `gorm:"size:32" json:"userType"`
}
