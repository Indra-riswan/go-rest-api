package main

import (
	"github.com/Indra-riswan/go-rest-api/config"
	"github.com/Indra-riswan/go-rest-api/controller"
	"github.com/Indra-riswan/go-rest-api/repository"
	"github.com/Indra-riswan/go-rest-api/srevice"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db            *gorm.DB                  = config.SetupConnectionDb()
	userRepsitory repository.UserRepository = repository.NewUserRepository(db)
	jwtService    srevice.JwtService        = srevice.NewJwtService()
	authService   srevice.AuthService       = srevice.NewAuthService(userRepsitory)
	authControl   controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer config.CloseConnectionDatabase(db)
	r := gin.Default()
	authRoutes := r.Group("api")
	{
		authRoutes.POST("/login", authControl.Login)
		authRoutes.POST("/register", authControl.Register)
	}

	r.Run()

}
