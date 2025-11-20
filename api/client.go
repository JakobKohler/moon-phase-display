package api

import (
	"net/http"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		BaseURL:    "https://api.visibleplanets.dev/v3",
		HTTPClient: &http.Client{},
	}
}
