package ai

import "context"

type AI interface {
	Talk(ctx context.Context, prompt string, content string) (string, error)
}
