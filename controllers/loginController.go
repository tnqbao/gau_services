package controllers

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tnqbao/gau_services/controllers/encrypt"
	"github.com/tnqbao/gau_services/models"
	"gorm.io/gorm"
)

type ClaimsResponse struct {
	UserID     uint   `json:"user_id"`
	Permission string `json:"permission"`
	jwt.RegisteredClaims
}

func getSecret(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening secret file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading secret file: %v", err)
	}
	return ""
}

var req RequestReceive

func Authentication(c *gin.Context) {
	jwtKey := getSecret("/run/secrets/jwt_key")

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("UserRequest binding error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserRequest binding error: " + err.Error()})
		return
	}
	if (req.Username == nil || req.Password == nil) && req.ExternalToken == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Either Username and Password or ExternalToken must be provided"})
		return
	}

	if *req.Username != "" && *req.Password != "" {
		hashedPassword := encrypt.HashPassword(*req.Password)

		user, err := verifyCredentials(c, *req.Username, hashedPassword)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid username or password!",
			})
			return
		}

		expirationTime := time.Now().Add(24 * time.Hour)
		claimsResponse := &ClaimsResponse{
			UserID:     user.UserId,
			Permission: user.Permission,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsResponse)
		fmt.Println([]byte(jwtKey))
		tokenString, err := token.SignedString([]byte(jwtKey))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		fmt.Println("generated JWT token:", tokenString)

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Username/Password or ExternalToken required"})
}

func verifyCredentials(c *gin.Context, username, password string) (models.User, error) {
	var user models.User
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Table("user_authentications").
		Select("users.user_id, users.permission").
		Joins("inner join users on users.user_id = user_authentications.user_id").
		Where("user_authentications.username = ? AND user_authentications.password = ?", username, password).
		First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
