package models

type UserInformation struct {
	UserId      uint    `gorm:"primaryKey" json:"user_id"`
	Name        *string `json:"name"`
	Phone       *string `gorm:"unique" json:"phone"`
	Email       *string `gorm:"unique" json:"email"`
	DateOfBirth *string `gorm:"unique" json:"date_of_birth"`
}
