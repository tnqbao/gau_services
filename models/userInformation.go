package models

import (
	"gorm.io/gorm"
)

type userInformaiton struct {
	gorm.Model
	userId uint   `gorm:"unique" json:"user_id"`
	Name   string `json:"name"`
	Phone  string ` json:"phone"`
}
