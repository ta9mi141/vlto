package api

import (
	"net/http"
)

const (
	BasicAuthPassword          string = "api_token"
	UserAgent                  string = "vlto"
	TogglSummaryReportEndPoint string = "https://toggl.com/reports/api/v2/summary"
)

type TogglReportsApiClient struct {
	Client            *http.Client
	ApiToken          string
	WorkSpaceId       string
	BasicAuthPassword string
	UserAgent         string
	EndPoint          string
}
