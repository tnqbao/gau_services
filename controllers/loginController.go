package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tnqbao/gau_services/controllers/encrypt"
	"github.com/tnqbao/gau_services/models"
	"gorm.io/gorm"
)

type ClaimsSend struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

type ClaimsReceive struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	ExternalToken string `json:"external_token"`
}

func Authentication(c *gin.Context) {
	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Secret key not set"})
		return
	}
	jwtService := encrypt.NewJWTService(jwtKey)

	fmt.Println("Secret Key:", jwtKey)

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is required"})
		return
	}
	tokenString := authHeader[len("Bearer "):]
	fmt.Println("Received Token:", tokenString)

	var claims ClaimsReceive
	err := jwtService.DecodeFromJWT(tokenString, &claims)
	if err != nil {
		fmt.Println("Token Decode Error:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	fmt.Println("Decoded Claims:", claims)

	if claims.Username != "" && claims.Password != "" {
		hashedPassword := encrypt.HashPassword(claims.Password)

		user, err := verifyCredentials(c, claims.Username, hashedPassword)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		expirationTime := time.Now().Add(24 * time.Hour)
		claimsSend := &ClaimsSend{
			UserID: user.UserId,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsSend)
		tokenString, err := token.SignedString([]byte(jwtKey))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		fmt.Println("Generated Token:", tokenString)

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Username/Password or ExternalToken required"})
}

func verifyCredentials(c *gin.Context, username, password string) (models.UserAuthentication, error) {
	var user models.UserAuthentication
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
