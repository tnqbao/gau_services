package models

import "time"

type UserInformation struct {
	UserId      uint       `gorm:"primaryKey" json:"user_id"`
	FullName    *string    `json:"fullname"`
	Phone       *string    `gorm:"unique" json:"phone"`
	Email       *string    `gorm:"unique" json:"email"`
	DateOfBirth *time.Time `json:"date_of_birth"`
}
