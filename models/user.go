package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	userId uint `gorm:"primaryKey" json:user_id`
}
