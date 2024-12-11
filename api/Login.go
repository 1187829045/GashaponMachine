package api

import (
	"GaMachine/form"
	"GaMachine/internal"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	user := form.User{}

	// 绑定请求中的数据到 user 结构体
	err := c.ShouldBind(&user)
	if err != nil {
		// 创建一个自定义错误，并返回客户端
		c.JSON(400, gin.H{
			"error": "绑定请求数据失败: " + err.Error(),
		})
		return // 直接返回，防止继续执行后续代码
	}
	fmt.Println(user.Username)
	// 这里可以继续进行登录逻辑，比如验证用户信息等
	// 假设验证成功，可以返回成功响应
	if isAdd, _ := internal.CheckNameInFile(internal.NameFile, user.Username); isAdd {
		c.JSON(200, gin.H{
			"message": "登录成功",
		})
		return
	}
	err = internal.WriteNamesToFile(internal.NameFile, []string{user.Username})
	if err != nil {
		c.JSON(400, gin.H{
			"error": "将用户名写入文件失败: " + err.Error(),
		})
	}

	DiamondCount := 1000

	err = internal.WriteDiamondCount(internal.BrickworkFile, user.Username, DiamondCount)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "将砖石数量写入文件失败: " + err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"message": "登录成功",
	})
}
