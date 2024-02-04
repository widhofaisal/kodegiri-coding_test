package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	Fullname   string `json:"fullname" form:"fullname" gorm:"not null"`
	Username   string `json:"username" form:"username" gorm:"not null, unique"`
	Password   string `json:"password" form:"password" gorm:"not null"`
	Role       string `json:"role" form:"role" gorm:"not null"`
}