package model

import (
	"GaMachine/consts"
	"GaMachine/global"
	"errors"
	"time"
)

type User struct {
	BaseModel
	UserId   uint       `gorm:"primaryKey;autoIncrement"`
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	UserName string     `gorm:"type:varchar(20)"`
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female表示女, male表示男'"`
}

func GetUser(mobile string) (user User, err error) {

	result := global.DB.Where("mobile = ?", mobile).Find(&user)

	if result.RowsAffected == 0 {
		return User{}, errors.New(consts.UserNotFound)
	}
	return user, nil
}

func CreateUser(user User) (err error) {

	result := global.DB.Create(&user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
