package api

import (
	"net/http"
)

const (
	basicAuthPassword          string = "api_token"
	userAgent                  string = "vlto"
	togglSummaryReportEndPoint string = "https://toggl.com/reports/api/v2/summary"
)

type togglReportsApiClient struct {
	client            *http.Client
	apiToken          string
	basicAuthPassword string
	userAgent         string
	workSpaceId       string
}
