package router

import (
	"GaMachine/internal/app/register/controller"
	"github.com/gin-gonic/gin"
)

func SetRegisterRouter(engin *gin.RouterGroup) {

	register := engin.Group("v1")
	{
		register.POST("/register", controller.Register)
	}
}
