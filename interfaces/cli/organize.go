package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var organizeCmd = &cobra.Command{
	Use:   "organize",
	Short: "Organize a directory to specification",
	Long:  `Organize a directory according to a given specification document`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("not implemented yet")
	},
}

func init() {
	fiorCmd.AddCommand(organizeCmd)
}
