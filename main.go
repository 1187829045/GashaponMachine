package main

import (
	"GaMachine/cmd"
	"GaMachine/initialization"
)

func init() {
	Init.HyInit()
	Init.GiftInit()
	Init.MysqlInit()
}
func main() {
	cmd.Execute()

}
