package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int32     `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

type User struct {
	BaseModel
	UserId   uint       `gorm:"primaryKey;autoIncrement"`
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	UserName string     `gorm:"type:varchar(20)"`
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female表示女, male表示男'"`
}

type Prize struct {
	BaseModel
	UserId   uint     `gorm:"type:int(11)"`
	Mobile   string   `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	UserName string   `gorm:"type:varchar(20)"`
	Prizes   []string `gorm:"type:varchar(20)"`
}

//type UpdateUserForm struct {
//	//Mobile   string `form:"mobile" json:"mobile" binding:"required"`
//	Name     string `form:"name" json:"name" binding:"required,min=3,max=10"`
//	Gender   string `form:"gender" json:"gender" binding:"required,oneof=female male"`
//	Birthday string `form:"birthday" json:"birthday" binding:"required,datetime=2006-01-02"`
//}
