package api

import (
	"net/http"
)

const (
	basicAuthPassword          string = "api_token"
	togglSummaryReportEndPoint string = "https://toggl.com/reports/api/v2/summary"
	userAgent                  string = "vlto"
)

type togglReportsApiClient struct {
	client            *http.Client
	apiToken          string
	basicAuthPassword string
	userAgent         string
	workSpaceId       string
}
