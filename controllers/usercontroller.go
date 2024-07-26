package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau_services/models"
	"gorm.io/gorm"
)

// add user
func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	var userInfor models.UserInformaiton
	var userAuth models.UserAuthentication

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&userInfor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&userAuth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		userInfor.UserId = user.UserId
		if err := tx.Create(&userInfor).Error; err != nil {
			return err
		}

		userAuth.UserId = user.UserId
		if err := tx.Create(&userAuth).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo người dùng"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Người dùng đã được tạo thành công"})
}
