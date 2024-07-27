package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau_services/models"
	"gorm.io/gorm"
)

// CREATE USER
func CreateUser(c *gin.Context, r Request) {
	db := c.MustGet("db").(*gorm.DB)

	userAuth := models.UserAuthentication{
		Username:      r.Username,
		Password:      r.Password,
		ExternalToken: r.ExternalToken,
	}
	userInfor := models.UserInformation{}

	err := db.Transaction(func(tx *gorm.DB) error {
		user := models.User{}
		var tokenSource *string
		if r.Username != nil {
			tokenSource = r.Username
		} else {
			tokenSource = r.ExternalToken
		}
		if tokenSource != nil {
			user.UserToken = generateToken(*tokenSource)
		} else {
			user.UserToken = generateToken("default_token_source")
		}
		userAuth.UserId = user.UserId
		userInfor.UserId = user.UserId

		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		if err := tx.Create(&userAuth).Error; err != nil {
			return err
		}
		if err := tx.Create(&userInfor).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Println("Transaction error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully created"})
}

// GET USER
type UserResponse struct {
	UserId      uint   `json:"user_id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
}

func GetUserById(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var user models.User
	var userInfo models.UserInformation

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user, "user_id = ?", id).Error; err != nil {
			return err
		}
		if err := tx.First(&userInfo, "user_id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	response := UserResponse{
		UserId:      user.UserId,
		Name:        toString(userInfo.Name),
		Phone:       toString(userInfo.Phone),
		Email:       toString(userInfo.Email),
		DateOfBirth: toString(userInfo.DateOfBirth),
	}

	c.JSON(http.StatusOK, response)
}

func generateToken(para string) string {
	h := md5.New()
	h.Write([]byte(strings.ToLower(para)))
	return hex.EncodeToString(h.Sum(nil))
}

func toString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
