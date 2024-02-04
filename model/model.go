package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model    `json:"-"`
	Email         string `json:"email" form:"email" gorm:"not null; unique"`
	Username      string `json:"username" form:"username" gorm:"not null; unique"`
	Fullname      string `json:"fullname" form:"fullname" gorm:"not null"`
	Password      string `json:"password" form:"password" gorm:"not null"`
	Phone         string `json:"phone" form:"phone" gorm:"not null"`
	BirthDate     string `json:"birth_date" form:"birth_date" gorm:"not null"`
	Address       string `json:"address" form:"address" gorm:"not null"`
	EarnedPoint   string `json:"earned_point" form:"earned_point" gorm:"not null"`
	RemainedPoint string `json:"remained_point" form:"remained_point" gorm:"not null"`
	RedeemedPoint string `json:"reedemed_point" form:"reedemed_point" gorm:"not null"`
	Referral      string `json:"referral" form:"referral" gorm:"not null"`
	Status        bool   `json:"bool" form:"bool" gorm:"not null"`
}
