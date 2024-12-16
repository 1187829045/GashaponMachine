package server

import (
	login_router "GaMachine/internal/app/Login/router"
	lottery_router "GaMachine/internal/app/lottery/router"
	query_router "GaMachine/internal/app/query/router"
	register_router "GaMachine/internal/app/register/router"
	"fmt"
	"net/http"
	"time"
	// adminrouter "go-server/internal/app/admin/router
	"github.com/gin-gonic/gin"
)

// NewHTTPRouter create http router.
func NewHTTPRouter() *gin.Engine {
	engine := gin.New()
	// TODO 默认需要注释掉pprof,可按需开启
	//pprof.Register(engine)

	fmt.Println("sever lottery in ", time.Now().Format(time.DateTime))
	// engine.Static("/bgo", "./background/")

	// middlewares.
	//engine.Use(middleware.LoggerConfig())
	//engine.Use(gzip.DefaultHandler().Gin)
	engine.Use(gin.Recovery())
	//engine.Use(middleware.Cors())
	//engine.Use(middleware.XRequest())
	//engine.Use(middleware.Recovery())
	// 加载404错误页面
	engine.NoRoute(func(c *gin.Context) {
		// 实现内部重定向
		c.JSON(http.StatusNotFound, gin.H{
			"title": "404 not found",
		})
	})
	// router group.
	app := engine.Group("/gash")
	{
		register_router.SetRegisterRouter(app)
		login_router.SetLoginRouter(app)
		lottery_router.SetLotteryRouter(app)
		query_router.SetQueryRouter(app)

	}
	return engine
}
