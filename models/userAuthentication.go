package models

type UserAuthentication struct {
	UserId   uint    `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username *string `gorm:"unique" json:"username"`
	Password *string `gorm:"size:255" json:"password"`
}
