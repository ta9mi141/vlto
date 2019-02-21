package config

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const config = `
ApiToken = "0123456789abcdefghijklmnopqrstuv"
WorkSpaceId = "1234567"
`

func TestInitDefaultConfigFilePath(t *testing.T) {
	defaultConfigFilePath := strings.Replace(
		defaultConfigFilePath, "$HOME", os.Getenv("HOME"), -1,
	)

	// If there already exists config file, rename it before test
	if _, err := os.Stat(defaultConfigFilePath); err == nil {
		os.Rename(defaultConfigFilePath, defaultConfigFilePath+".tmp")
		defer os.Rename(defaultConfigFilePath+".tmp", defaultConfigFilePath)
	}

	err := ioutil.WriteFile(defaultConfigFilePath, []byte(config), 0644)
	defer os.Remove(defaultConfigFilePath)
	if err != nil {
		t.Fatal(err)
	}

	if err := Init(""); err != nil {
		t.Error(err)
	}

	if viper.GetString("ApiToken") != "0123456789abcdefghijklmnopqrstuv" {
		t.Error(err)
	}

	if viper.GetString("WorkSpaceId") != "1234567" {
		t.Error(err)
	}
}

func TestInitCustomConfigFilePath(t *testing.T) {
	customConfigFilePath := "vlto.toml"
	err := ioutil.WriteFile(customConfigFilePath, []byte(config), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(customConfigFilePath)

	if err := Init(customConfigFilePath); err != nil {
		t.Error(err)
	}

	if viper.GetString("ApiToken") != "0123456789abcdefghijklmnopqrstuv" {
		t.Error(err)
	}

	if viper.GetString("WorkSpaceId") != "1234567" {
		t.Error(err)
	}
}
