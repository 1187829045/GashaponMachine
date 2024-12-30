package router

import (
	"GaMachine/internal/app/gachaSystem/controller"
	"GaMachine/middlewares"
	"github.com/gin-gonic/gin"
)

func SetLotteryRouter(engin *gin.RouterGroup) {

	lottery := engin.Group("v1").Use(middlewares.JWTAuth())
	{
		lottery.POST("/lottery", controller.Lottery)
		lottery.GET("/prize", controller.GetPrize)
	}
}
