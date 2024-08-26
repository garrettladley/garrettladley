package operating_cash

import (
	"context"
	"fmt"
	"net/http"

	"github.com/garrettladley/garrettladley/subs/reports/pkg/dts/request"
)

func query(ctx context.Context, params params) (resp response, err error) {
	requestURL, err := params.build()
	if err != nil {
		return resp, fmt.Errorf("failed to build request URL: %w", err)
	}

	err = request.WithContext(ctx, http.MethodGet, requestURL, nil, &resp)

	return
}
