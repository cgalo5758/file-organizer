package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {

	var cmdVersion = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of file-organizer",
		Long:  `Print the version number of file-organizer in semantic version format.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("file-organizer version 0.0.0")
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdVersion)
	rootCmd.Execute()
}
