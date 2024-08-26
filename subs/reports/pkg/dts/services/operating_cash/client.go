package operating_cash

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/garrettladley/left-arrow/subs/reports/pkg/dts/constants"
	"github.com/garrettladley/left-arrow/subs/reports/pkg/dts/utilities"
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

func (p *params) build() (*url.URL, error) {
	queryParams := utilities.URLValues{}

	if len(p.Fields) > 0 {
		queryParams.Add("fields", strings.Join(p.Fields, ","))
	}

	if len(p.Filters) > 0 {
		filterStrings := make([]string, len(p.Filters))
		for i, f := range p.Filters {
			if len(f.Value) == 0 {
				continue
			}
			filterStrings[i] = f.string()
		}
		if len(filterStrings) > 0 {
			queryParams.Add("filter", strings.Join(filterStrings, ","))
		}
	}

	if len(p.Sorts) > 0 {
		sortStrings := make([]string, len(p.Sorts))
		for i, s := range p.Sorts {
			if len(s.Field) == 0 {
				continue
			}
			sortStrings[i] = s.string()
		}
		if len(sortStrings) > 0 {
			queryParams.Add("sort", strings.Join(sortStrings, ","))
		}
	}

	if p.PageSize > 0 {
		queryParams.Add("page[size]", strconv.Itoa(p.PageSize))
	}

	if p.PageNumber > 0 {
		queryParams.Add("page[number]", strconv.Itoa(p.PageNumber))
	}

	requestURL, err := url.Parse(fmt.Sprintf("%s%s", constants.BASE_URL, constants.DEPOSITS_AND_WITHDRAWALS_OF_OPERATING_CASH_URL))
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	requestURL.RawQuery = queryParams.Encode()

	return requestURL, nil
}

type result struct {
	response response
	err      error
}
