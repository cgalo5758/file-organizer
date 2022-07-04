package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of fior",
	Long:  "Print the version number of fior in semantic version format.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fior version 0.0.0")
	},
}

func init() {
	fiorCmd.AddCommand(versionCmd)
}
