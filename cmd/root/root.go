package root

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "root",
	Short: "Gashapon machine entrance",
	Long:  `Root command for the gashapon machine entrance`,
}

func Execute() {
	RootCmd.Execute()
}
