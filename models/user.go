package models

type User struct {
	UserId     uint   `gorm:"primaryKey" json:"user_id"`
	Permission string `json:"user_permission"`
}
