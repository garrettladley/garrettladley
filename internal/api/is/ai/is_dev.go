//go:build !prod
// +build !prod

package ai

import (
	"context"
	"fmt"

	"github.com/garrettladley/garrettladley/internal/api/is/predicate"
)

func (c *Client) Is(ctx context.Context, n int64, p predicate.Predicate) (bool, error) {
	switch p {
	case predicate.Even:
		return n%2 == 0, nil
	case predicate.Odd:
		return n%2 != 0, nil
	default:
		return false, fmt.Errorf("unsupported predicate: %v", p)
	}
}
