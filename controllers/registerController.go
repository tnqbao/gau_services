package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau_services/controllers/encrypt"
)

type Request struct {
	Username      *string `json:"Username"`
	Password      *string `json:"Password"`
	ExternalToken *string `json:"ExternalToken"`
}

func Register(c *gin.Context) {
	var req Request

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("UserRequest binding error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserRequest binding error: " + err.Error()})
		return
	}
	*req.Password = encrypt.HashPassword(*req.Password)
	if (req.Username == nil || req.Password == nil) && req.ExternalToken == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Either Username and Password or ExternalToken must be provided"})
		return
	}

	log.Println("Parsed Request:", req)

	CreateUser(c, req)
}
