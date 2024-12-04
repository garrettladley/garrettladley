package openai

import (
	"github.com/sashabaranov/go-openai"
)

type Client struct {
	client openai.Client
	cfg    Config
}

type Config struct {
	Model string
}

func New(key string, cfgs ...Config) *Client {
	return &Client{
		client: *openai.NewClient(key),
		cfg:    configure(cfgs...),
	}
}

func configure(cfgs ...Config) Config {
	if len(cfgs) < 1 {
		return DefaultConfig
	}
	cfg := cfgs[0]
	if cfg.Model == "" {
		cfg.Model = DefaultConfig.Model
	}
	return cfg
}

var DefaultConfig = Config{
	Model: openai.GPT3Dot5Turbo,
}
