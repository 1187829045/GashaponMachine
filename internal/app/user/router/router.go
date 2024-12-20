package router

import (
	"GaMachine/internal/app/user/controller"
	"github.com/gin-gonic/gin"
)

func SetUserRouter(engin *gin.RouterGroup) {
	api := engin.Group("v1")
	{
		api.POST("login", controller.Login)
		api.POST("/register", controller.Register)
	}

}
