package ai

import (
	"github.com/garrettladley/garrettladley/pkg/ai"
)

type Client struct {
	ai ai.AI
}

func New(ai ai.AI) *Client {
	return &Client{
		ai: ai,
	}
}
