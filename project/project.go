package project

import (
	"errors"
	"github.com/it-akumi/toggl-go/reports"
	"github.com/spf13/viper"
	"time"
)

type config struct {
	Name          string
	TargetHour    int
	StartDate     time.Time
	IterationDays int
}

type projectStatus struct {
	name                 string
	targetHour           int
	totalAchievedSec     int
	iterationAchievedSec int
	lastDate             string
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

func fetchAchievedSec(projectName string, span dateSpan) (int, error) {
	client := reports.NewClient(viper.GetString("apiToken"))
	resp, err := client.GetSummary(&reports.RequestParameters{
		UserAgent:   "vlto",
		WorkSpaceId: viper.GetString("workSpaceId"),
		Since:       span.since,
		Until:       span.until,
	})
	if err != nil {
		return 0, err
	}
	for _, datum := range resp.Data {
		if datum.Title.Project == projectName {
			return datum.Time / 1000, nil // Time entries are in milliseconds
		}
	}
	return 0, nil
}

func estimateLastDate(unachievedSec, iterationAchievedSec, iterationDays int, now time.Time) (string, error) {
	if iterationAchievedSec < 0 || iterationDays <= 0 {
		return "", errors.New("Invalid iterationAchievedSec or iterationDays")
	}
	if unachievedSec <= 0 {
		return "Finished", nil
	}
	if iterationAchievedSec == 0 {
		return "Never", nil
	}

	// Round up unachievedSec / iterationAchievedSec
	remainingDays := (unachievedSec + iterationAchievedSec - 1) / iterationAchievedSec * iterationDays
	return now.AddDate(0, 0, remainingDays).Format("2006-01-02"), nil
}
