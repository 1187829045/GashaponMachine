package router

import (
	"GaMachine/internal/app/query/controller"
	"GaMachine/middlewares"
	"github.com/gin-gonic/gin"
)

func SetQueryRouter(engin *gin.RouterGroup) {

	query := engin.Group("v1")
	{
		query.GET("/query", middlewares.JWTAuth(), controller.QueryPrize)
	}
}
