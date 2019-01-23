package project

import (
	"time"
)

type Project struct {
	Name      string
	Goal      int
	StartDate time.Time
	Iteration int
}
