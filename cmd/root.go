package cmd

import (
	"fmt"
	"github.com/it-akumi/toggl-go/reports"
	"github.com/it-akumi/vlto/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// flags
var cfgFilePath string

var rootCmd = &cobra.Command{
	Use:     "vlto",
	Short:   "vlto shows velocity of your projects of Toggl",
	Version: "0.0",
	Run: func(cmd *cobra.Command, args []string) {
		client := reports.NewClient(viper.GetString("apiToken"))
		fmt.Printf("%+v\n", client)
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
}

func initConfig() {
	config.Init(cfgFilePath)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
