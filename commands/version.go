package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of file-organizer",
	Long:  "Print the version number of file-organizer in semantic version format.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("file-organizer version 0.0.0")
	},
}

func init() {
	filorgCmd.AddCommand(versionCmd)
}
