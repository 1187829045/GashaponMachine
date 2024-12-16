package form

type User struct {
	Username string `json:"username" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PassWordLoginForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"` //手机号码格式有规范可寻， 自定义validator
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
	UserName string `form:"username" json:"username" binding:"required"`
}

type RegisterForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
}
type RequestPrize struct {
	UserName string `json:"username" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
}
