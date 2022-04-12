// Cobra should have a package alias before this code becomes a zoo.
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

// Execute adds all child commands to the root command and sets flags appropriately.
//
// Wait... Is this supposed to be capitalized?
// I thought this gets capitalized once exported.
func Execute() {
	cobra.CheckErr(fiorCmd.Execute())
}
