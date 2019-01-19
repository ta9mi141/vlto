package cmd

import (
	"fmt"
	"github.com/it-akumi/vlto/config"
	"github.com/it-akumi/vlto/toggl"
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
		client, err := toggl.NewReportsClient(
			viper.GetString("apiToken"),
			viper.GetString("workSpaceId"),
			"vlto",
			toggl.SummaryReportEndPoint,
		)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("%v\n", client)
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
