package project

import (
	"github.com/spf13/viper"
	"time"
)

type config struct {
	Name          string
	Target        int
	StartDate     time.Time
	IterationDays int
}

type project struct {
	name                  string
	target                int
	totalAchievedHour     int
	iterationAchievedHour int
	lastDay               time.Time
}

func Unmarshal() ([]config, error) {
	var configs []config
	if err := viper.UnmarshalKey("Projects", &configs); err != nil {
		return configs, err
	}
	return configs, nil
}
