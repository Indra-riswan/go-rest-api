package controller

import "github.com/gin-gonic/gin"

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	//service yang dibutuhkan
}

func NewAuthController() AuthController {
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hallo Login",
	})
}

func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hallo Register",
	})
}
