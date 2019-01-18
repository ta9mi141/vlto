package toggl

import (
	"net/http"
	"net/url"
)

const (
	basicAuthPassword     string = "api_token" // Defined in Toggl Reports API v2
	userAgent             string = "vlto"
	SummaryReportEndPoint string = "https://toggl.com/reports/api/v2/summary"
)

type ReportsApiClient struct {
	Client      *http.Client
	ApiToken    string
	WorkSpaceId string
	UserAgent   string
	Url         *url.URL
}

type ReportsApiError struct {
	Message    string `json:"message"`
	Tip        string `json:"tip"`
	StatusCode int    `json:"code"`
}

type SummaryReport struct {
	Error *ReportsApiError `json:"error,omitempty`
	Data  []struct {
		Title struct {
			Project string `json:"project"`
		} `json:"title"`
		Time int `json:"time"`
	} `json:"data"`
}
