package models

type User struct {
	UserId     uint   `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Permission string `json:"user_permission"`
}
