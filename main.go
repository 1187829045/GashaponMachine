package main

import (
	"GaMachine/cmd"
	"GaMachine/initialization/giftInit"
	"GaMachine/initialization/hystrixInit"
)

func init() {
	hystrixInit.Init()
	giftInit.Init()
}
func main() {
	cmd.Execute()

}
