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

func getIterationSpan(now time.Time, iterationDays int) dateSpan {
	iterationStartDate := now.AddDate(0, 0, -iterationDays)
	return dateSpan{since: iterationStartDate, until: now}
}

func divideElapsedYears(startDate, now time.Time) []dateSpan {
	elapsedYears := make([]dateSpan, 0)
	for {
		oneYearLaterFromStart := startDate.AddDate(1, 0, 0)
		if now.After(oneYearLaterFromStart) {
			elapsedYears = append(elapsedYears, dateSpan{
				since: startDate,
				until: oneYearLaterFromStart,
			})
			startDate = oneYearLaterFromStart.AddDate(0, 0, 1)
		} else {
			elapsedYears = append(elapsedYears, dateSpan{
				since: startDate,
				until: now,
			})
			return elapsedYears
		}
	}
}
