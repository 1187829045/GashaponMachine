package controller

import (
	"GaMachine/global"
	"GaMachine/middlewares"
	"GaMachine/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func QueryPrize(c *gin.Context) {

	claims, _ := c.Get("claims")                      // 从上下文中获取 "claims" 信息
	currentUser := claims.(*middlewares.CustomClaims) // 将 "claims" 转换为自定义的用户声明类型

	Prize := model.Prize{
		UserId: currentUser.ID,
	}

	result := global.DB.Find(&Prize)
	if result.RowsAffected == 0 {
		c.JSON(200, gin.H{
			"message": "没有中奖信息",
		})
		return
	}
	fmt.Println(Prize.Prizes)

	c.JSON(200, gin.H{
		"name":  currentUser.NickName,
		"prize": Prize.Prizes,
	})

}
