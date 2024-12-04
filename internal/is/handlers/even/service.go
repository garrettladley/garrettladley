package even

import (
	is "github.com/garrettladley/garrettladley/internal/is/ai/openai"
	"github.com/garrettladley/garrettladley/pkg/ai/openai"
)

type Service struct {
	client *is.Client
}

func newService(client *openai.Client) *Service {
	return &Service{
		client: is.New(client),
	}
}
