package main

import (
	"github.com/Indra-riswan/go-rest-api/config"
	"github.com/Indra-riswan/go-rest-api/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db          *gorm.DB = config.SetupConnectionDb()
	authControl          = controller.NewAuthController()
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
