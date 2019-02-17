package config

import (
	"github.com/spf13/viper"
)

func Init(cfgFilePath string) error {
	viper.AddConfigPath("$HOME/.config") // Adding $HOME/.config as first search path
	viper.SetConfigName("vlto")          // Name of config file (without extention)
	if cfgFilePath != "" {
		viper.SetConfigFile(cfgFilePath)
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
