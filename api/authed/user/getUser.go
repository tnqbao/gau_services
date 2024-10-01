package api_user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	provider "github.com/tnqbao/gau_services/api"
	"github.com/tnqbao/gau_services/models"
	"gorm.io/gorm"
)

func GetUserById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	token, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userToken := token.(*provider.UserToken)

	userPermission, exists := c.Get("user_permission")
	if !exists {
		userPermission = "member"
	}

	userPermissionStr, ok := userPermission.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid permission format"})
		return
	}

	if userToken.UserId != id && userPermissionStr != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

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
	response := provider.ServerResp{
		UserId:     user.UserId,
		Fullame:    provider.ToString(userInfo.Fullname),
		Email:      provider.ToString(userInfo.Email),
		DateBirth:  provider.ToString(userInfo.DateOfBirth),
		Permission: user.Permission,
	}
	c.Set("user_permission", user.Permission)
	c.JSON(http.StatusOK, response)
}
