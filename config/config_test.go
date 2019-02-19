package config

import (
	"io/ioutil"
	"os"
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

func TestInitCustomConfigFilePath(t *testing.T) {
	customConfigFilePath := "vlto.toml"
	ioutil.WriteFile(customConfigFilePath, []byte(config), 0644)
	defer os.Remove(customConfigFilePath)

	if err := Init(customConfigFilePath); err != nil {
		t.Error()
	}
}
