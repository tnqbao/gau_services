package models

type UserInformation struct {
	UserId      uint    `gorm:"primaryKey" json:"user_id"`
	Fullname    *string `json:"fullname"`
	Phone       *string `gorm:"unique" json:"phone"`
	Email       *string `gorm:"unique" json:"email"`
	DateOfBirth *string `gorm:"unique" json:"date_of_birth"`
}
