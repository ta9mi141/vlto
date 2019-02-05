package project

import (
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
