package models

import (
	"gorm.io/gorm"
)

type UserAuthentication struct {
	gorm.Model
	UserId       uint   `gorm:"primaryKey" json:"user_id"`
	Username     string `gorm:"unique"`
	Password     string `gorm:"password"`
	ExtenalToken string `gorm:"extenal_token"`
}
