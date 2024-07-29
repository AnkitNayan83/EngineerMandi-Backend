package controllers

import (
	"net/http"

	"github.com/AnkitNayan83/EngineerMandi-Backend/services"
	"github.com/AnkitNayan83/EngineerMandi-Backend/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ctrl *AuthController) GoogleLogin(c *gin.Context) {
	url := ctrl.authService.GetGoogleLoginUrl()

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (ctrl *AuthController) GoogleCallback(c *gin.Context) {
	code := c.Query("code")

	token, err := ctrl.authService.ExchangeCodeForToken(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userInfo, err := ctrl.authService.FetchUserInfo(token)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, isNewUser, err := ctrl.authService.HandleUserLogin(userInfo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwtToken, err := utils.GenerateJwt(newUser.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken, "isNewUser": isNewUser})
}
