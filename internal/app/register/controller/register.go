package controller

import (
	"GaMachine/form"
	"GaMachine/global"
	"GaMachine/model"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	registerForm := form.RegisterForm{}

	if err := c.ShouldBind(&registerForm); err != nil {
		c.JSON(400, gin.H{
			"error": "参数错误",
		})
		return
	}
	result := global.DB.Find(&model.User{
		Mobile: registerForm.Mobile,
	})
	if result.RowsAffected != 0 {
		c.JSON(400, gin.H{
			"error": "该手机号已注册",
		})
	}

	user := model.User{
		UserName: registerForm.Username,
		Mobile:   registerForm.Mobile,
		Password: registerForm.PassWord,
	}
	result = global.DB.Create(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "注册失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}
