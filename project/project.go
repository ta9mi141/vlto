package project

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/it-akumi/toggl-go/reports"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
)

type config struct {
	Name          string
	TargetHour    int
	StartDate     time.Time
	IterationDays int
}

type status struct {
	Name                  string
	TargetHour            string
	TotalAchievedHour     string
	IterationAchievedHour string
	LastDate              string
}

type dateSpan struct {
	since, until time.Time
}

type summaryReport struct {
	Data []struct {
		Title struct {
			Project string `json:"project"`
		} `json:"title"`
		Time int `json:"time"`
	} `json:"data"`
}

func unmarshal() ([]config, error) {
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

func fetchAchievedSec(projectName string, span dateSpan, achievedSecChan chan<- int, errorChan chan<- error) {
	client := reports.NewClient(viper.GetString("apiToken"))
	summaryReport := new(summaryReport)

	// operation is retried using Exponential Backoff and Jitter
	// when client.GetSummary fails due to '429 Too Many Requests'
	operation := func() error {
		err := client.GetSummary(
			context.Background(),
			&reports.SummaryRequestParameters{
				StandardRequestParameters: &reports.StandardRequestParameters{
					UserAgent:   "vlto",
					WorkSpaceId: viper.GetString("workSpaceId"),
					Since:       span.since,
					Until:       span.until,
				},
			},
			summaryReport,
		)

		if err != nil {
			var reportsError reports.Error
			if errors.As(err, &reportsError) && reportsError.StatusCode() == http.StatusTooManyRequests {
				return err
			} else {
				// operation should not be retried when client.GetSummary returns
				// an error except '429 Too Many Requests'
				return backoff.Permanent(err)
			}
		}
		return nil
	}
	if err := backoff.Retry(operation, backoff.NewExponentialBackOff()); err != nil {
		errorChan <- err
		return
	}
	for _, datum := range summaryReport.Data {
		if datum.Title.Project == projectName {
			achievedSecChan <- datum.Time / 1000 // Time entries are in milliseconds
			return
		}
	}
	achievedSecChan <- 0
	return
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

func generateStatus(c config, statusChan chan<- *status, errorChan chan<- error) {
	achievedSecChan := make(chan int)
	fetchErrorChan := make(chan error)

	elapsedYears := divideElapsedYears(c.StartDate, time.Now())
	for _, year := range elapsedYears {
		go fetchAchievedSec(c.Name, year, achievedSecChan, fetchErrorChan)
	}

	totalAchievedSec := 0
	for i := 0; i < len(elapsedYears); i++ {
		select {
		case achievedSec := <-achievedSecChan:
			totalAchievedSec += achievedSec
		case err := <-fetchErrorChan:
			errorChan <- err
			return
		}
	}

	go fetchAchievedSec(
		c.Name,
		getIterationSpan(time.Now(), c.IterationDays),
		achievedSecChan,
		fetchErrorChan,
	)

	iterationAchievedSec := 0
	select {
	case achievedSec := <-achievedSecChan:
		iterationAchievedSec = achievedSec
	case err := <-fetchErrorChan:
		errorChan <- err
		return
	}

	lastDate, err := estimateLastDate(
		c.TargetHour*3600-totalAchievedSec,
		iterationAchievedSec,
		c.IterationDays,
		time.Now(),
	)
	if err != nil {
		errorChan <- err
		return
	}

	statusChan <- &status{
		Name:                  c.Name,
		TargetHour:            fmt.Sprintf("%d", c.TargetHour),
		TotalAchievedHour:     fmt.Sprintf("%.1f", float64(totalAchievedSec)/3600),
		IterationAchievedHour: fmt.Sprintf("%.1f", float64(iterationAchievedSec)/3600),
		LastDate:              lastDate,
	}
	return
}

const (
	TableFormat string = "table"
	JsonFormat  string = "json"
)

func toTable(projectsStatus []status) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Target", "Total", "Iteration", "LastDate"})
	for _, status := range projectsStatus {
		table.Append([]string{
			status.Name,
			status.TargetHour,
			status.TotalAchievedHour,
			status.IterationAchievedHour,
			status.LastDate,
		})
	}
	table.Render()
}

func toJSON(projectsStatus []status) error {
	output, err := json.Marshal(projectsStatus)
	if err != nil {
		return err
	}
	os.Stdout.Write(output)
	return nil
}

func Show(format string) error {
	if !(format == "" || format == TableFormat || format == JsonFormat) {
		return fmt.Errorf(`Valid format is "%s" or "%s"`, TableFormat, JsonFormat)
	}

	projectsConfig, err := unmarshal()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	statusChan := make(chan *status)
	errorChan := make(chan error)
	for _, config := range projectsConfig {
		go generateStatus(config, statusChan, errorChan)
	}

	projectsStatus := []status{}
	for i := 0; i < len(projectsConfig); i++ {
		select {
		case status := <-statusChan:
			projectsStatus = append(projectsStatus, *status)
		case err := <-errorChan:
			return err
		}
	}

	switch format {
	case TableFormat:
		toTable(projectsStatus)
	case JsonFormat:
		if err := toJSON(projectsStatus); err != nil {
			return err
		}
	default:
		toTable(projectsStatus)
	}

	return nil
}
