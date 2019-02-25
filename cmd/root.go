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

var rootCmd = &cobra.Command{
	Use:     "vlto",
	Short:   "vlto shows velocity of your projects of Toggl",
	Version: "0.0",
	Run: func(cmd *cobra.Command, args []string) {
		projectConfigs, err := project.Unmarshal()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, config := range projectConfigs {
			status, err := project.GenerateProjectStatus(&config)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("%+v\n", status)
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
