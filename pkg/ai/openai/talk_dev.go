//go:build !prod
// +build !prod

package openai

import (
	"context"
)

func (c *Client) Talk(ctx context.Context, prompt string, content string) (string, error) {
	return "recieved the following | prompt: [" + prompt + "] | content: [" + content + "]", nil
}
