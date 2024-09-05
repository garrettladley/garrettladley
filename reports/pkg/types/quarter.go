package types

import (
	"log/slog"
	"time"
)

type Quarter string

const (
	Quarter1 Quarter = "Q1"
	Quarter2 Quarter = "Q2"
	Quarter3 Quarter = "Q3"
	Quarter4 Quarter = "Q4"
)

// QuarterFromTime returns the quarter from the given time.
func QuarterFromTime(t time.Time) Quarter {
	month := t.Month()
	switch {
	case month >= time.January && month <= time.March:
		return Quarter1
	case month >= time.April && month <= time.June:
		return Quarter2
	case month >= time.July && month <= time.September:
		return Quarter3
	case month >= time.October && month <= time.December:
		return Quarter4
	default:
		slog.Warn("unknown quarter", "month", month)
		return ""
	}
}

// QuarterFromString returns the quarter from the given string.
func QuarterFromString(s string) Quarter {
	switch s {
	case "Q1":
		return Quarter1
	case "Q2":
		return Quarter2
	case "Q3":
		return Quarter3
	case "Q4":
		return Quarter4
	default:
		slog.Warn("unknown quarter", "s", s)
		return ""
	}
}

// Int returns the integer representation of the quarter.
func (q Quarter) Int() int {
	switch q {
	case Quarter1:
		return 1
	case Quarter2:
		return 2
	case Quarter3:
		return 3
	case Quarter4:
		return 4
	default:
		slog.Warn("unknown quarter", "q", q)
		return 0
	}
}
