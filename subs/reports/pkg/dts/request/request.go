package request

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	go_json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

// Executes the request with context and stores the result
// in the value pointed to by v
func WithContext(ctx context.Context, method string, url *url.URL, body io.Reader, v interface{}) error {
	var r requester
	switch method {
	case http.MethodGet:
		r = fiber.Get
	case http.MethodHead:
		r = fiber.Head
	case http.MethodPost:
		r = fiber.Post
	case http.MethodPut:
		r = fiber.Put
	case http.MethodPatch:
		r = fiber.Patch
	case http.MethodDelete:
		r = fiber.Delete
	default:
		return fmt.Errorf("unsupported method: %s", method)
	}

	agent := r(url.String()).
		JSONEncoder(go_json.Marshal).
		JSONDecoder(go_json.Unmarshal)

	if body != nil {
		bodyBytes, err := io.ReadAll(body)
		if err != nil {
			return err
		}

		agent = agent.Body(bodyBytes)
	}

	errCh := make(chan error)

	go func() {
		defer close(errCh)
		statusCode, _, errs := agent.Struct(&v)
		if len(errs) > 0 {
			errCh <- fmt.Errorf("failed to make request: %w", errors.Join(errs...))
			return
		}

		if statusCode != http.StatusOK {
			errCh <- fmt.Errorf("received status code: %d", statusCode)
			return
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err, ok := <-errCh:
			if !ok {
				return nil
			}
			return err
		}
	}
}

type requester func(url string) *fiber.Agent
