package project

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Name          string
	Goal          int
	StartDate     time.Time
	IterationDays int
}

type Project struct {
	Name                  string
	Goal                  int
	TotalAchievedHour     int
	IterationAchievedHour int
	LastDay               time.Time
}

func Unmarshal() []Config {
	var configs []Config
	viper.UnmarshalKey("Projects", &configs)
	return configs
}
