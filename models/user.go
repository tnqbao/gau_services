package models

type User struct {
	UserId    uint   `gorm:"primaryKey;autoIncrement" json:"user_id"`
	UserToken string `gorm:"unique" json:"user_token"`
}
