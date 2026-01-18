package stream

import (
	"fmt"

	streamsdk "github.com/GetStream/stream-chat-go/v5"
)

type Client struct {
	client *streamsdk.Client
}

type Config struct {
	APIKey    string
	APISecret string
}

func New(cfg Config) (*Client, error) {
	c, err := streamsdk.NewClient(cfg.APIKey, cfg.APISecret)
	if err != nil {
		return nil, fmt.Errorf("init getstream client: %w", err)
	}

	return &Client{client: c}, nil
}
