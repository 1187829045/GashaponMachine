package service

import (
	"GaMachine/consts"
	"GaMachine/internal/app/user/dto"
	"GaMachine/middlewares"
	"GaMachine/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type IUser interface {
	Register(c *gin.Context) (err error)
	Login(c *gin.Context) (token string, err error)
}
type User struct{}

func NewUser() IUser {
	return &User{}
}
func (u *User) Register(c *gin.Context) error {

	registerForm := dto.RegisterForm{}

	if err := c.ShouldBind(&registerForm); err != nil {
		err = errors.New(consts.ErrInvalidParameter)
		return err
	}
	userInfo, err := model.GetUser(registerForm.Mobile)

	if err != nil {
		err = errors.New(consts.UserExist)
		return err
	}

	if err = model.CreateUser(userInfo); err != nil {
		err = errors.New("注册失败")
		return err
	}
	return nil
}

func (user *User) Login(c *gin.Context) (token string, err error) {

	login_form := dto.PassWordLoginForm{}

	// 绑定请求中的数据到 user 结构体
	err = c.ShouldBind(&login_form)
	if err != nil {
		// 创建一个自定义错误，并返回客户端
		err = errors.New(consts.ErrInvalidParameter)
		return
	}
	userInfo, err := model.GetUser(login_form.Mobile)

	if err != nil {
		err = errors.New(consts.UserExist)
		return
	}

	j := middlewares.NewJWT()
	claims := middlewares.CustomClaims{
		ID:       userInfo.UserId,
		NickName: login_form.UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               //签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
			Issuer:    "llb",
		},
	}
	token, err = j.CreateToken(claims)
	if err != nil {
		err = errors.New("生成token失败")
		return
	}
	return
}
