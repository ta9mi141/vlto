package cmd

import (
	"fmt"
	"github.com/it-akumi/vlto/api"
	"github.com/it-akumi/vlto/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

// flags
var cfgFile string

var rootCmd = &cobra.Command{
	Use:     "vlto",
	Short:   "vlto shows velocity of your projects of Toggl",
	Version: "0.0",
	Run: func(cmd *cobra.Command, args []string) {
		client := &api.TogglReportsApiClient{
			Client:            &http.Client{},
			ApiToken:          viper.GetString("apiToken"),
			WorkSpaceId:       viper.GetString("WorkSpaceId"),
			BasicAuthPassword: api.BasicAuthPassword,
			UserAgent:         api.UserAgent,
			EndPoint:          api.TogglSummaryReportEndPoint,
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
