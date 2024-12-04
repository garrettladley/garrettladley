package openai

import (
	"context"
	"fmt"
	"strconv"

	"github.com/garrettladley/garrettladley/internal/is/predicate"
)

func (c *Client) Is(ctx context.Context, n int64, p predicate.Predicate) (bool, error) {
	resp, err := c.client.Talk(
		ctx,
		predicate.IntoPrompt(p),
		strconv.FormatInt(n, 10),
	)
	if err != nil {
		return false, err
	}
	is, err := strconv.ParseBool(resp)
	if err != nil {
		return false, fmt.Errorf("recieved unexpected response from OpenAI: %v", resp)
	}
	return is, nil
}
