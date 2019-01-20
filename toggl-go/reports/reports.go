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
	client            *http.Client
	basicAuthPassword string
	apiToken          string
	workSpaceId       string
	userAgent         string
	url               *url.URL
}

func NewClient(apiToken, workSpaceId, userAgent, endpoint string) (*client, error) {
	if len(apiToken) == 0 {
		return nil, errors.New("Missing API token")
	}
	if len(workSpaceId) == 0 {
		return nil, errors.New("Missing workspace id")
	}
	if len(userAgent) == 0 {
		return nil, errors.New("Missing user agent")
	}
	if len(endpoint) == 0 {
		return nil, errors.New("Missing end point")
	}
	url, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	newClient := &client{
		client:            &http.Client{},
		basicAuthPassword: basicAuthPassword,
		apiToken:          apiToken,
		workSpaceId:       workSpaceId,
		userAgent:         userAgent,
		url:               url,
	}
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
