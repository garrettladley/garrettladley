package operating_cash

import (
	"context"
	"fmt"
	"time"
)

func Query(ctx context.Context, start time.Time, end time.Time, fields ...string) (resp Response, err error) {
	res, err := query(ctx, params{Filters: []filter{after(start), before(end), xfields(fields...)}})
	if err != nil {
		return resp, fmt.Errorf("failed to query operating cash: %w", err)
	}

	resp = into(res)
	return
}
