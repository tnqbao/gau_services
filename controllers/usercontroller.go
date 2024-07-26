package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau_services/models"
	"gorm.io/gorm"
)

type UserRequest struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	DateOfBirth   string `json:"date_of_birth"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	ExternalToken string `json:"external_token"`
}

func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var req UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("UserRequest binding error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserRequest binding error: " + err.Error()})
		return
	}
	userInfo := models.UserInformaiton{
		Name:        req.Name,
		Email:       req.Email,
		Phone:       req.Phone,
		DateOfBirth: req.DateOfBirth,
	}

	userAuth := models.UserAuthentication{
		Username:     req.Username,
		Password:     req.Password,
		ExtenalToken: req.ExternalToken,
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		user := models.User{
			UserToken: generateToken(),
		}
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		userInfo.UserId = user.ID
		userAuth.UserId = user.ID
		if err := tx.Create(&userInfo).Error; err != nil {
			return err
		}
		if err := tx.Create(&userAuth).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Println("Transaction error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo người dùng: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Người dùng đã được tạo thành công"})
}
func generateToken() string {
	return "random-token"
}
