package router

import (
	"GaMachine/internal/app/query/controller"
	"github.com/gin-gonic/gin"
)

func SetQueryRouter(engin *gin.RouterGroup) {

	query := engin.Group("v1")
	{
		query.GET("/query", controller.QueryPrize)
	}
}
