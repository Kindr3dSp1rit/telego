package telego

import (
	"context"
	"net/http"
	"telego/pkg/api"
)

const (
	DefaultAPIHost = "https://api.telegram.org"
)

type Bot struct {
	apiCli *api.ClientWithResponses

	cfg       *config
	authToken string
}

func NewBotClient(authToken string, opts ...Option) (*Bot, error) {
	cfg := &config{
		apiHost: DefaultAPIHost,
		httpCli: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	apiCli, err := api.NewClientWithResponses(cfg.apiHost, api.WithHTTPClient(cfg.httpCli))
	if err != nil {
		return nil, err
	}

	return &Bot{
		apiCli:    apiCli,
		cfg:       cfg,
		authToken: authToken,
	}, nil
}

func (b *Bot) GetMe(ctx context.Context) (*api.User, error) {
	resp, err := b.apiCli.GetBotTokenGetMeWithResponse(ctx, b.authToken)
	if err != nil {
		return nil, err
	}
	return &resp.JSON200.Result, nil
}
