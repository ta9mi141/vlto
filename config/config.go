package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func Init(cfgFilePath string) {
	viper.AddConfigPath("$HOME/.config") // Adding $HOME/.config as first search path
	viper.SetConfigName("vlto")          // Name of config file (without extention)
	if cfgFilePath != "" {
		viper.SetConfigFile(cfgFilePath)
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
