package main

import (
	"GaMachine/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/login", api.Login)

	r.Run(":8080")
}
