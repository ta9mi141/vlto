package cmd

import (
	"fmt"
	"os"

	"github.com/it-akumi/vlto/config"
	"github.com/it-akumi/vlto/project"
	"github.com/spf13/cobra"
)

// flags
var cfgFilePath string
var format string
var httpsProxy string

var rootCmd = &cobra.Command{
	Use:     "vlto",
	Short:   "vlto shows velocity of your projects of Toggl",
	Version: "0.0",
	Run: func(cmd *cobra.Command, args []string) {
		// --proxy option overrides an environment variable HTTPS_PROXY.
		// Since a child process cannot change environment variables of its parent process,
		// it's not necessary to rewrite it to original value.
		if err := os.Setenv("HTTPS_PROXY", httpsProxy); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
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
	rootCmd.PersistentFlags().StringVar(
		&httpsProxy,
		"proxy",
		"",
		"the URL of an environment variable HTTPS_PROXY",
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
