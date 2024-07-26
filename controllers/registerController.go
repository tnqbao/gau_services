package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	if (req.Username == nil || req.Password == nil) && req.ExternalToken == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Either Username and Password or ExternalToken must be provided"})
		return
	}

	log.Println("Parsed Request:", req)

	CreateUser(c, req)
}
