package models

type UserAuthentication struct {
	UserId        uint    `gorm:"primaryKey" json:"user_id"`
	Username      *string `gorm:"unique" json:"username"`
	Password      *string `gorm:"size:255" json:"password"`
	ExternalToken *string `gorm:"size:255" json:"external_token"`
}
