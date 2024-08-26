package operating_cash

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	go_json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func query(ctx context.Context, params params) (response, error) {
	requestUrl, err := params.build()
	if err != nil {
		return response{}, fmt.Errorf("failed to build request URL: %w", err)
	}

	agent := fiber.Get(requestUrl.String())
	agent.JSONDecoder(go_json.Unmarshal)

	resultCh := make(chan result)

	go func() {
		var resp response
		statusCode, _, errs := agent.Struct(&resp)
		if len(errs) > 0 {
			resultCh <- result{response{}, fmt.Errorf("failed to make request: %w", errors.Join(errs...))}
			close(resultCh)
			return
		}

		if statusCode != http.StatusOK {
			resultCh <- result{response{}, fmt.Errorf("received status code: %d", statusCode)}
			close(resultCh)
			return
		}

		resultCh <- result{resp, nil}
		close(resultCh)
	}()

	for {
		select {
		case <-ctx.Done():
			return response{}, ctx.Err()
		case res := <-resultCh:
			return res.response, res.err
		}
	}
}

type result struct {
	response response
	err      error
}
