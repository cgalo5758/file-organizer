package cli

import (
	"github.com/cgalo5758/fior/internal"
	"github.com/spf13/cobra"
)

var organizeCmd = &cobra.Command{
	Use:   "organize",
	Short: "Organize a directory to specification",
	Long:  `Organize a directory according to a given specification document`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		internal.Organize(args[0], args[1])
	},
}

func init() {
	fiorCmd.AddCommand(organizeCmd)
}
