package api_user

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	provider "github.com/tnqbao/gau_services/api"
	"github.com/tnqbao/gau_services/models"
	"gorm.io/gorm"
)

func Authentication(c *gin.Context) {
	var req models.UserAuthentication
	jwtKey := os.Getenv("JWT_SECRET")
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("UserRequest binding error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}
	if req.Username == nil || req.Password == nil || *req.Username == "" || *req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and Password are required"})
		return
	}
	hashedPassword := provider.HashPassword(*req.Password)
	user, err := verifyCredentials(c, *req.Username, hashedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &provider.ClaimsResponse{
		UserID: user.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	c.SetCookie("auth_token", tokenString, 3600*24, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func verifyCredentials(c *gin.Context, username, password string) (models.User, error) {
	var user models.User
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Table("user_authentications").
		Select("users.user_id, users.permission").
		Joins("INNER JOIN users ON users.user_id = user_authentications.user_id").
		Where("user_authentications.username = ? AND user_authentications.password = ?", username, password).
		First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}