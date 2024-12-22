package model

import (
	"GaMachine/consts"
	"GaMachine/global"
	"errors"
)

type Prize struct {
	BaseModel
	UserId   uint     `gorm:"type:int(11)"`
	Mobile   string   `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	UserName string   `gorm:"type:varchar(20)"`
	Prizes   []string `gorm:"type:varchar(20)"`
}

func GetPrize(userId uint) ([]string, error) {
	var prizes Prize

	result := global.DB.Where("user_id = ?", userId).Find(&prizes)

	if result.RowsAffected == 0 {
		return prizes.Prizes, errors.New(consts.UserNotFound)
	}
	return prizes.Prizes, nil
}

func AddPrize(userId uint, prizes []string) error {
	p, err := GetPrize(userId)
	if err != nil {
		return errors.New("查询商品失败")
	}
	p = append(p, prizes...)
	err = CreatePrize(userId, p)
	if err != nil {
		return err
	}
	return nil
}

func CreatePrize(userId uint, prizes []string) error {
	prize := Prize{}
	prize.Prizes = prizes
	result := global.DB.Save(&prize)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
