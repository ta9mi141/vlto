package config

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const config = `
ApiToken = "0123456789abcdefghijklmnopqrstuv"
WorkSpaceId = "1234567"

[[Projects]]
Name = "Sample Project 1"
Target = 1000
StartDate = 2016-10-11T00:00:00+00:00
IterationDays = 7

[[Projects]]
Name = "Sample Project 2"
Target = 2000
StartDate = 2019-01-01T00:00:00+00:00
IterationDays = 14
`

func TestInitDefaultConfigFilePath(t *testing.T) {
	defaultConfigFilePath := strings.Replace(
		defaultConfigFilePath, "$HOME", os.Getenv("HOME"), -1,
	)
	err := ioutil.WriteFile(defaultConfigFilePath, []byte(config), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(defaultConfigFilePath)

	if err := Init(""); err != nil {
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
}
