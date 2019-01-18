package api

import (
	"net/http"
	"net/url"
)

const (
	basicAuthPassword          string = "api_token" // Defined in Toggl Reports API v2
	userAgent                  string = "vlto"
	TogglSummaryReportEndPoint string = "https://toggl.com/reports/api/v2/summary"
)

type TogglReportsApiClient struct {
	Client      *http.Client
	ApiToken    string
	WorkSpaceId string
	UserAgent   string
	Url         *url.URL
}

type TogglReportsApiError struct {
	Message    string `json:"message"`
	Tip        string `json:"tip"`
	StatusCode int    `json:"code"`
}

type TogglSummaryReport struct {
	Error *TogglReportsApiError `json:"error,omitempty`
	Data  []struct {
		Title struct {
			Project string `json:"project"`
		} `json:"title"`
		Time int `json:"time"`
	} `json:"data"`
}
