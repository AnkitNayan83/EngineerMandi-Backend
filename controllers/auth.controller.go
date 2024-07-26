package controllers

import (
	"net/http"

	"github.com/AnkitNayan83/SMA-backend/services"
	"github.com/AnkitNayan83/SMA-backend/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService services.UserService
}

func NewAuthController(userService services.UserService) *AuthController {
	return &AuthController{
		userService: userService,
	}
}

func (ctrl *AuthController) GoogleLogin(c *gin.Context) {
	url := ctrl.userService.GetGoogleLoginUrl()

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (ctrl *AuthController) GoogleCallback(c *gin.Context) {
	code := c.Query("code")

	token, err := ctrl.userService.ExchangeCodeForToken(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userInfo, err := ctrl.userService.FetchUserInfo(token)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := ctrl.userService.HandleUserLogin(userInfo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwtToken, err := utils.GenerateJwt(newUser.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}
