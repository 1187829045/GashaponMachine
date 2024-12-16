package controller

import (
	"GaMachine/form"
	"GaMachine/global"
	"GaMachine/middlewares"
	"GaMachine/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	lFrom := form.PassWordLoginForm{}

	// 绑定请求中的数据到 user 结构体
	err := c.ShouldBind(&lFrom)
	if err != nil {
		// 创建一个自定义错误，并返回客户端
		c.JSON(400, gin.H{
			"error": "绑定请求数据失败: " + err.Error(),
		})
		return
	}
	result := global.DB.Find(&model.User{
		Mobile: lFrom.Mobile,
	})
	if result.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"error": "用户不存在",
		})
	}

	j := middlewares.NewJWT()
	claims := middlewares.CustomClaims{
		NickName: lFrom.UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               //签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
			Issuer:    "llb",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "登陆成功",
		"token":   token,
	})
}
