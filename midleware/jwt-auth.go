package midleware

import (
	"log"
	"net/http"

	"github.com/Indra-riswan/go-rest-api/helper"
	"github.com/Indra-riswan/go-rest-api/srevice"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizedJwt(jwtService srevice.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			respons := helper.BuildErrorsRespon("Failed to procces request ", "No token found,", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, respons)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[user_id] :", claims["user_id"])
			log.Println("Claims[issuer] :", claims["issuer"])
		} else {
			log.Println(err)
			respons := helper.BuildErrorsRespon("erorr wrong token", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
		}
	}
}
