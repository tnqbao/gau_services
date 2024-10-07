package api_user

import "github.com/golang-jwt/jwt/v5"

type ClientReq struct {
	Username    *string `json:"username"`
	Password    *string `json:"password"`
	Fullname    *string `json:"fullname"`
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
	DateOfBirth *string `json:"date_of_birth"`
}

type ServerResp struct {
	UserId     uint   `json:"user_id"`
	Fullame    string `json:"fullname"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	DateBirth  string `json:"date_of_birth"`
	Permission string `json:"permission"`
}

type UserToken struct {
	UserId string `json:"user_id"`
}

type ClaimsResponse struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}
