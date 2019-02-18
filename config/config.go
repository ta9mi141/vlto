package config

import (
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

const (
	defaultConfigFilePath string = "$HOME/.config/vlto.toml"
)

func Init(cfgFilePath string) error {
	// Adding a directory of defaultConfigFilePath as first search path
	viper.AddConfigPath(filepath.Dir(defaultConfigFilePath))

	// Name of config file (without extention)
	viper.SetConfigName(strings.TrimSuffix(
		filepath.Base(defaultConfigFilePath),
		filepath.Ext(defaultConfigFilePath),
	))
	if cfgFilePath != "" {
		viper.SetConfigFile(cfgFilePath)
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
