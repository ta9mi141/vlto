package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// flags
var cfgFile string

var rootCmd = &cobra.Command{
	Use:     "vlto",
	Short:   "vlto shows velocity of your projects of Toggl",
	Version: "0.0",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file (default is $HOME/.config/vlto.toml)",
	)
}

func initConfig() {
	viper.AddConfigPath("$HOME/.config") // Adding $HOME/.config as first search path
	viper.SetConfigName("vlto")          // Name of config file (without extention)

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	if err := viper.ReadInConfig(); err != nil {
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
