package feeder

import (
	"net/http"
	"time"
)

type Client struct {
	handler http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		handler: http.Client{
			Timeout: timeout,
		},
	}
}
