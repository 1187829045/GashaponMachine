package start

import (
	"GaMachine/cmd/root"
	"fmt"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")

	},
}

func init() {
	root.RootCmd.AddCommand(startCmd)
}
