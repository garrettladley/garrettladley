package openai

import "github.com/garrettladley/garrettladley/pkg/ai/openai"

type Client struct {
	client *openai.Client
}

func New(client *openai.Client) *Client {
	return &Client{
		client: client,
	}
}
