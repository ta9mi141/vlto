package cmd

import (
	"fmt"
	"github.com/it-akumi/toggl-go/reports"
	"github.com/it-akumi/vlto/config"
	"github.com/it-akumi/vlto/project"
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
		client := reports.NewSummaryClient(viper.GetString("apiToken"))
		fmt.Printf("%+v\n", client)
		var prj []project.Project
		viper.UnmarshalKey("Projects", &prj)
		fmt.Printf("%+v\n", prj)
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
	config.Init(cfgFile)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
