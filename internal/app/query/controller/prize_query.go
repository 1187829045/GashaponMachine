package controller

import (
	"GaMachine/internal/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

func QueryPrize(c *gin.Context) {
	username := c.GetHeader("username")

	fmt.Println(username)

	if isAdd, _ := common.CheckNameInFile(common.NameFile, username); !isAdd {
		c.JSON(400, gin.H{
			"error": "用户不存在",
		})
		return
	}
	prizes, err := common.ReadPrizes(common.GiftFile, username)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "读取奖品失败: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"prize": prizes,
	})
}
