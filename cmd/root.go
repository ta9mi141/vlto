package cmd

import (
	"fmt"
	"github.com/it-akumi/vlto/config"
	"github.com/it-akumi/vlto/project"
	"github.com/spf13/cobra"
	"os"
)

// flags
var cfgFilePath string
var format string

var rootCmd = &cobra.Command{
	Use:     "vlto",
	Short:   "vlto shows velocity of your projects of Toggl",
	Version: "0.0",
	Run: func(cmd *cobra.Command, args []string) {
		if err := project.Show(format); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(
		&cfgFilePath,
		"config",
		"",
		"config file (default is $HOME/.config/vlto.toml)",
	)
	rootCmd.PersistentFlags().StringVar(
		&format,
		"format",
		"",
		"the output format 'table' (default) or 'json'",
	)
}

func initConfig() {
	if err := config.Init(cfgFilePath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
