package api_user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserById(c *gin.Context) {
	// db := c.MustGet("db").(*gorm.DB)
	// id := c.Param("id")
	// token, exists := c.Get("user_id")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	// 	return
	// }
	// fmt.Print(token)
	// userToken := token.(*provider.UserToken)
	// userPermission, exists := c.Get("permission")
	// if !exists {
	// 	userPermission = "member"
	// }

	// if userToken.UserId != id && userPermission != "admin" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	// 	return
	// }
	// err := db.Transaction(func(tx *gorm.DB) error {
	// 	if result := tx.Delete(&models.User{}, id); result.Error != nil {
	// 		return result.Error
	// 	} else if result.RowsAffected == 0 {
	// 		return gorm.ErrRecordNotFound
	// 	}

	// 	if err := tx.Delete(&models.UserAuthentication{}, id).Error; err != nil {
	// 		return err
	// 	}
	// 	if err := tx.Delete(&models.UserInformation{}, id).Error; err != nil {
	// 		return err
	// 	}

	// 	return nil
	// })

	// if err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	// 	} else {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	}
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func DeleteListUser(c *gin.Context) {

}
