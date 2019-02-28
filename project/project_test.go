package project

import (
	"reflect"
	"testing"
	"time"
)

func TestGetIterationSpan(t *testing.T) {
	today := time.Date(2019, 2, 4, 0, 0, 0, 0, time.Local)
	twoWeeksAgo := time.Date(2019, 1, 21, 0, 0, 0, 0, time.Local)
	expected := dateSpan{
		since: twoWeeksAgo,
		until: today,
	}
	iterationSpan := getIterationSpan(today, 14)
	if iterationSpan != expected {
		t.Error()
	}
}

func TestDivideElapsedYears(t *testing.T) {
	today := time.Date(2019, 2, 5, 0, 0, 0, 0, time.Local)
	startDate := time.Date(2016, 10, 11, 0, 0, 0, 0, time.Local)
	expected := []dateSpan{
		{
			since: startDate,
			until: time.Date(2017, 10, 11, 0, 0, 0, 0, time.Local),
		},
		{
			since: time.Date(2017, 10, 12, 0, 0, 0, 0, time.Local),
			until: time.Date(2018, 10, 12, 0, 0, 0, 0, time.Local),
		},
		{
			since: time.Date(2018, 10, 13, 0, 0, 0, 0, time.Local),
			until: today,
		},
	}
	elapsedYears := divideElapsedYears(startDate, today)
	if !reflect.DeepEqual(elapsedYears, expected) {
		t.Error()
	}
}

func TestEstimateLastDate(t *testing.T) {
	today := time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)

	lastDate, err := estimateLastDate(100, 10, 2, today)
	if lastDate != "2019-01-21" {
		t.Error()
	}

	lastDate, err = estimateLastDate(0, 0, 2, today)
	if lastDate != "Finished" {
		t.Error()
	}

	lastDate, err = estimateLastDate(100, 0, 2, today)
	if lastDate != "Never" {
		t.Error()
	}

	lastDate, err = estimateLastDate(100, 0, -1, today)
	if err == nil {
		t.Error()
	}
}
