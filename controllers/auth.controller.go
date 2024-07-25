package controllers

import (
	"net/http"

	"github.com/AnkitNayan83/SMA-backend/services"
	"github.com/gin-gonic/gin"
)

func GoogleLogin(c *gin.Context) {
	url := services.GetGoogleLoginUrl()

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := services.ExchangeCodeForToken(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userInfo, err := services.FetchUserInfo(token)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userInfo)
}
