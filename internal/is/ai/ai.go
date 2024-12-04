package ai

import (
	"context"

	"github.com/garrettladley/garrettladley/internal/is/predicate"
)

type AI interface {
	Is(ctx context.Context, n int64, p predicate.Predicate) (bool, error)
}
