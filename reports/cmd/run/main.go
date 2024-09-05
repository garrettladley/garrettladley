package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/garrettladley/garrettladley/common/pkg/builder"
	"github.com/garrettladley/garrettladley/reports/pkg/services/consumer_credit"
)

func main() {
	start := builder.
		NewDate().
		Day(22).
		Month(time.August).
		Year(2024).
		Loc(time.UTC).
		MustBuild()

	end := builder.
		NewDate().
		Day(23).
		Month(time.August).
		Year(2024).
		Loc(time.UTC).
		MustBuild()

	// resp, err := operating_cash.Query(context.Background(), start, end)

	// slog.Info("query", "resp", resp, "err", err)

	data, err := consumer_credit.Query(context.Background(), start, end)

	slog.Info("consumer_credit", "data", data, "err", err)
}
