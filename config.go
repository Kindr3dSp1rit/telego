package telego

import "net/http"

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Option func(*config)

func WithAPIHost(apiHost string) Option {
	return func(c *config) {
		c.apiHost = apiHost
	}
}

func WithHTTPClient(client HTTPClient) Option {
	return func(c *config) {
		c.httpCli = client
	}
}

type config struct {
	apiHost string
	httpCli HTTPClient
}
