package api_user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	provider "github.com/tnqbao/gau_services/api"
	"github.com/tnqbao/gau_services/models"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context, r provider.ClientReq) {
	db := c.MustGet("db").(*gorm.DB)
	err := db.Transaction(func(tx *gorm.DB) error {
		user := models.User{}
		user.Permission = "member"
		userAuth := models.UserAuthentication{
			Username: r.Username,
			Password: r.Password,
			UserId:   user.UserId,
		}
		userInfor := models.UserInformation{
			Fullname:    r.Fullname,
			Email:       r.Email,
			DateOfBirth: provider.FormatStringToDate(r.DateOfBirth),
			UserId:      user.UserId,
		}

		if err := tx.Create(&userAuth).Error; err != nil {
			return err
		}
		if err := tx.Create(&user).Error; err != nil {
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
