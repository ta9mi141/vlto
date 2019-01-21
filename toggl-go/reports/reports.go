package reports

import (
	"errors"
	"net/http"
	"net/url"
)

const (
	basicAuthPassword string = "api_token" // Defined in Toggl Reports API v2
	SummaryEndpoint   string = "https://toggl.com/reports/api/v2/summary"
)

type client struct {
	httpClient *http.Client
	header     http.Header
	apiToken   string
	url        *url.URL
}

func NewClient(apiToken, endpoint string) (*client, error) {
	if len(apiToken) == 0 {
		return nil, errors.New("Missing API token")
	}
	if len(endpoint) == 0 {
		return nil, errors.New("Missing end point")
	}
	url, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	newClient := &client{
		httpClient: &http.Client{},
		header:     make(http.Header),
		apiToken:   apiToken,
		url:        url,
	}
	newClient.header.Set("Content-type", "application/json")

	return newClient, nil
}

type Error struct {
	Message    string `json:"message"`
	Tip        string `json:"tip"`
	StatusCode int    `json:"code"`
}

type Summary struct {
	Error *Error `json:"error,omitempty`
	Data  []struct {
		Title struct {
			Project string `json:"project"`
		} `json:"title"`
		Time int `json:"time"`
	} `json:"data"`
}
