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

type dateSpan struct {
	since, until time.Time
}

func Unmarshal() ([]config, error) {
	var configs []config
	if err := viper.UnmarshalKey("Projects", &configs); err != nil {
		return configs, err
	}
	return configs, nil
}

func getIterationSpan(today time.Time, iterationDays int) dateSpan {
	iterationStartDate := today.AddDate(0, 0, -iterationDays)
	return dateSpan{since: iterationStartDate, until: today}
}
