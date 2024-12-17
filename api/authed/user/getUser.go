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
	tokenId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userPermissionStr := "member"
	if userPermission, exists := c.Get("user_permission"); exists {
		if val, ok := userPermission.(string); ok {
			userPermissionStr = val
		}
	}

	tokenIdStr, ok := tokenId.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid id format"})
		return
	}

	if tokenIdStr != id {
		if userPermissionStr != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to access this resource!"})
			return
		}
	}

	var user models.UserAuthentication
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
		Fullame:    provider.ToString(userInfo.FullName),
		Email:      provider.ToString(userInfo.Email),
		DateBirth:  provider.FormatDateToString(userInfo.DateOfBirth),
		Permission: user.Permission,
	}
	c.JSON(http.StatusOK, response)
}
