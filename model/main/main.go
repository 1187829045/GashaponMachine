package main

import (
	"GaMachine/global"
	Init "GaMachine/initialization"
	"GaMachine/model"
)

func main() {
	Init.MysqlInit()

	_ = global.DB.AutoMigrate(&model.User{})
	_ = global.DB.AutoMigrate(&model.Prize{})
}
