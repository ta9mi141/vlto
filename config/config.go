package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"time"
)

type projectConfig struct {
	Name      string
	Goal      int
	StartDate time.Time
	Iteration int
}

type config struct {
	ApiToken    string
	WorkSpaceId string
	Projects    []projectConfig
}

func Init(cfgFile string) {
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
