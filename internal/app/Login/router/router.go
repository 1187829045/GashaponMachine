package router

import (
	"GaMachine/internal/app/Login/controller"
	"github.com/gin-gonic/gin"
)

func SetLoginRouter(engin *gin.RouterGroup) {
	login := engin.Group("v1")
	{
		login.POST("login", controller.Login)
	}
}
