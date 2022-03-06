package commands

import (
	"github.com/spf13/cobra"
)

// filorgCmd represents the base command when called without subcommands
var filorgCmd = &cobra.Command{
	Use:   "file-organizer",
	Short: "file-organizer is a tool to organize files and directories",
	Long: `file-organizer is a tool to organize files and directories
	based on a given specification.`,
}

func Execute() {
	cobra.CheckErr(filorgCmd.Execute())
}
