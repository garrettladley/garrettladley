package even

import (
	is "github.com/garrettladley/garrettladley/internal/api/is/ai"
	"github.com/garrettladley/garrettladley/pkg/ai"
)

type Service struct {
	client *is.Client
}

func New(ai ai.AI) *Service {
	return &Service{
		client: is.New(ai),
	}
}
