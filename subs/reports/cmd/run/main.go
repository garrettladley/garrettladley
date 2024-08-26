package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/garrettladley/garrettladley/subs/reports/pkg/dts/services/operating_cash"
)

func main() {
	start := time.Date(2024, time.August, 22, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, time.August, 23, 0, 0, 0, 0, time.UTC)

	resp, err := operating_cash.Query(context.Background(), start, end)

	slog.Info("query", "resp", resp, "err", err)

}
