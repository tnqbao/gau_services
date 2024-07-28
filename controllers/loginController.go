package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tnqbao/gau_services/models"
	"gorm.io/gorm"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

var Json struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	ExternalToken string `json:"externalToken"`
}

func Authentication(c *gin.Context) {

	if err := c.ShouldBindJSON(&Json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.UserAuthentication
	var err error

	// if Json.ExternalToken != "" {
	// 	// Process external token (e.g., Google, Facebook)
	// 	user, err = verifyExternalToken(c, Json.ExternalToken)
	// 	if err != nil {
	// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid external token"})
	// 		return
	// 	}
	// } else
	if Json.Username != "" && Json.Password != "" {
		user, err = verifyCredentials(c, Json.Username, Json.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username/Password or ExternalToken required"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// func verifyExternalToken(c *gin.Context, token string) (models.UserAuthentication, error) {
// 	var user models.UserAuthentication
// 	if token == "valid_external_token" {
// 		username := "externalUser"
// 		user.Username = &username
// 		return user, nil
// 	}
// 	return user, errors.New("invalid external token")
// }

func verifyCredentials(c *gin.Context, username, password string) (models.UserAuthentication, error) {
	var user models.UserAuthentication
	db := c.MustGet("db").(*gorm.DB)
	hashedPassword := hashPassword(password)
	if err := db.Where("username = ? AND password = ?", username, hashedPassword).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func hashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}
