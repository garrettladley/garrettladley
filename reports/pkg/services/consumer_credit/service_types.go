package consumer_credit

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Data struct {
	StudentLoans      TimeSeries
	MotorVehicleLoans TimeSeries
}

type TimeSeries struct {
	Date Date
	Data float32 // in millions
}

type Date struct {
	Quarter Quarter
	Year    int
}

type Quarter string

const (
	Quarter1 Quarter = "Q1"
	Quarter2 Quarter = "Q2"
	Quarter3 Quarter = "Q3"
	Quarter4 Quarter = "Q4"
)

func TimeSeriesFrom(s string) (studentLoans TimeSeries, motorVehicleLoans TimeSeries) {
	date := dateFrom(s)
	studentLoansData, endIdx := parseData(s, 25)
	next := readUntil(s, endIdx, func(r byte) bool { return unicode.IsDigit(rune(r)) })
	motorVehicleLoansData, _ := parseData(s, next)

	studentLoans = TimeSeries{
		Date: date,
		Data: studentLoansData,
	}

	motorVehicleLoans = TimeSeries{
		Date: date,
		Data: motorVehicleLoansData,
	}

	return
}

func parseData(s string, startIdx int) (data float32, endIdx int) {
	var builder strings.Builder
	n := len(s)

	for idx := startIdx; idx < n; idx++ {
		r := s[idx]
		switch r {
		case ',':
			// Skip commas
		case '.':
			builder.WriteByte(r)
		case '\n':
			return parseFloat(builder.String()), idx
		default:
			if unicode.IsDigit(rune(r)) {
				builder.WriteByte(r)
			} else {
				slog.Warn("unexpected character", "r", fmt.Sprintf("'%c'", r))
			}
		}
	}

	return 0.0, -1
}

func readUntil(s string, start int, predicate func(byte) bool) int {
	n := len(s)
	for idx := start; idx < n; idx++ {
		if predicate(s[idx]) {
			return idx
		}
	}
	return n
}

func parseFloat(str string) float32 {
	data, err := strconv.ParseFloat(str, 32)
	if err != nil {
		slog.Warn("error parsing data", "err", err, "str", str)
		return 0.0
	}
	return float32(data)
}

func dateFrom(s string) Date {
	year, err := strconv.Atoi(s[9:13])
	if err != nil {
		slog.Warn("error parsing year", "err", err, "s", s)
	}
	quarter := quarterFrom(s[14:16])
	return Date{
		Year:    year,
		Quarter: quarter,
	}
}

func quarterFrom(s string) Quarter {
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

func whatQuarter(now time.Time) Quarter {
	month := now.Month()
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

func (q Quarter) into() int {
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
