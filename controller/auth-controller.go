package controller

import (
	"net/http"
	"strconv"

	"github.com/Indra-riswan/go-rest-api/dto"
	"github.com/Indra-riswan/go-rest-api/entity"
	"github.com/Indra-riswan/go-rest-api/helper"
	"github.com/Indra-riswan/go-rest-api/srevice"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authservice srevice.AuthService
	jwtservice  srevice.JwtService
}

func NewAuthController(auth srevice.AuthService, jwt srevice.JwtService) AuthController {
	return &authController{
		authservice: auth,
		jwtservice:  jwt,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDto dto.LoginDto
	errDTO := ctx.ShouldBind(&loginDto)
	if errDTO != nil {
		respons := helper.BuildErrorsRespon("Failed to proces request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, respons)
		return
	}
	authservice := c.authservice.VerifyCredintial(loginDto.Email, loginDto.Password)
	if v, ok := authservice.(entity.User); ok {
		generatetoken := c.jwtservice.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatetoken
		respons := helper.BuildRespons(true, "Ok!", v)
		ctx.JSON(200, respons)
		return
	}
	respons := helper.BuildErrorsRespon("Failed to sent Request", "invalid token", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, respons)

}

func (c *authController) Register(ctx *gin.Context) {
	var registerDto dto.RegisterDto
	errDto := ctx.ShouldBind(&registerDto)
	if errDto != nil {
		respons := helper.BuildErrorsRespon("Failed to send Request", errDto.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respons)
		return
	}
	if !c.authservice.IsduplicatEmail(registerDto.Email) {
		respons := helper.BuildErrorsRespon("Failed to sent Request", "Ducplicat Email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, respons)
	} else {
		usercrated := c.authservice.CreateUser(registerDto)
		token := c.jwtservice.GenerateToken(strconv.FormatUint(usercrated.ID, 10))
		usercrated.Token = token
		respons := helper.BuildRespons(true, "Ok!", usercrated)
		ctx.JSON(http.StatusCreated, respons)
	}

}
