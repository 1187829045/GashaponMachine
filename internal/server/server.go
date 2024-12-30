package server

import (
	lottery_router "GaMachine/internal/app/gachaSystem/router"

	login_router "GaMachine/internal/app/user/router"
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
		login_router.SetUserRouter(app)
		lottery_router.SetLotteryRouter(app)

	}
	return engine
}
