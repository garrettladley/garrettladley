package operating_cash

import (
	"context"
	"fmt"
	"time"
)

func Query(ctx context.Context, start time.Time, end time.Time, fields ...string) (Response, error) {
	res, err := query(ctx, params{Filters: []filter{after(start), before(end), xfields(fields...)}})
	if err != nil {
		return Response{}, fmt.Errorf("failed to query operating cash: %w", err)
	}

	return into(res), nil
}
