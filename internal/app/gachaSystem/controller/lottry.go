package controller

import (
	"GaMachine/internal/app/gachaSystem/dto"
	"GaMachine/internal/app/gachaSystem/service"
	"github.com/gin-gonic/gin"
)

func Lottery(c *gin.Context) {
	var req dto.Lottry
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	gaSystem := service.NewGachaSystem()
	gaSystem.Lottery(c, req)

}
