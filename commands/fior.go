package commands

import (
	"github.com/spf13/cobra"
)

// fiorCmd represents the base command when called without subcommands
var fiorCmd = &cobra.Command{
	Use:   "fior",
	Short: "fior is a tool to organize files and directories",
	Long: `fior is a tool to organize files and directories
	based on a given specification.`,
}

func Execute() {
	cobra.CheckErr(fiorCmd.Execute())
}
