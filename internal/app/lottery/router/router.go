package router

import (
	"GaMachine/internal/app/lottery/controller"
	"github.com/gin-gonic/gin"
)

func SetLotteryRouter(engin *gin.RouterGroup) {

	lottery := engin.Group("v1")
	{
		lottery.POST("/lottery", controller.Lottery)
	}
}
