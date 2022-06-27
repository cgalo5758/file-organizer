package commands

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	cli "github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// fiorCmd represents the base command when called without subcommands
var fiorCmd = &cli.Command{
	Use:   "fior",
	Short: "fior is a tool to organize files and directories",
	Long: `fior is a tool to organize files and directories
	based on a given specification.`,
}

func init() {
	cli.OnInitialize(initConfig)
	fiorCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	fiorCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	fiorCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	fiorCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	fiorCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cli.CheckErr(fiorCmd.Execute())
}
