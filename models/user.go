package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId    uint   `gorm:"primaryKey" json:"user_id"`
	UserToken string `gorm:"unique" json:"user_token"`
}
