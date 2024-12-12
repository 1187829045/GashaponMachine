package main

import (
	"GaMachine/cmd"
	"GaMachine/initialization"
)

func init() {

	initialization.InitPrize()
}
func main() {
	cmd.Execute()

}
