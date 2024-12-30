package controller

import (
	"GaMachine/internal/app/gachaSystem/service"
	"github.com/gin-gonic/gin"
)

func GetPrize(c *gin.Context) {
	gaSystem := service.NewGachaSystem()
	gaSystem.GetPrize(c)
}
