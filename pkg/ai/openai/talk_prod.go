//go:build prod
// +build prod

package openai

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func (c *Client) Talk(ctx context.Context, prompt string, content string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: c.cfg.Model,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: prompt},
			{Role: openai.ChatMessageRoleUser, Content: content},
		},
	}
	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
