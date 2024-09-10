package consumer_credit

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"unicode"

	"github.com/garrettladley/garrettladley/internal/reports/types"
)

type parser struct {
	s    string
	n    int
	read int
}

func parse(s string) Data {
	p := parser{s: s, n: len(s)}
	return p.parseData()
}

func (p parser) parseData() Data {
	qDate := parseDate(p.s)

	studentLoansData := p.parseFloat(25)

	motorVehicleStart := p.readUntil(p.s, p.read, func(r byte) bool { return unicode.IsDigit(rune(r)) })
	motorVehicleLoansData := p.parseFloat(motorVehicleStart)

	return Data{
		StudentLoans: types.TimeSeries{
			QuarterDate: qDate,
			Data:        studentLoansData,
		},
		MotorVehicleLoans: types.TimeSeries{
			QuarterDate: qDate,
			Data:        motorVehicleLoansData,
		},
	}
}

func (p *parser) parseFloat(start int) float32 {
	var builder strings.Builder

	for idx := start; idx < p.n; idx++ {
		r := p.s[idx]
		switch r {
		case ',':
			// Skip commas
		case '.':
			builder.WriteByte(r)
		case '\n':
			p.read = idx
			return parseFloat(builder.String())
		default:
			if unicode.IsDigit(rune(r)) {
				builder.WriteByte(r)
			} else {
				slog.Warn("unexpected character", "r", fmt.Sprintf("'%c'", r))
			}
		}
	}

	return 0.0
}
func (p *parser) readUntil(s string, startIdx int, pred func(r byte) bool) (endIdx int) {
	for idx := startIdx; idx < p.n; idx++ {
		if pred(s[idx]) {
			p.read = idx
			return idx
		}
	}
	return p.n
}

func parseFloat(s string) float32 {
	data, err := strconv.ParseFloat(s, 32)
	if err != nil {
		slog.Warn("error parsing data", "err", err, "s", s)
		return 0.0
	}
	return float32(data)
}

func parseDate(s string) types.QuarterDate {
	year, err := strconv.Atoi(s[9:13])
	if err != nil {
		slog.Warn("error parsing year", "err", err, "s", s)
	}

	quarter := types.QuarterFromString(s[14:16])

	return types.QuarterDate{
		Year:    year,
		Quarter: quarter,
	}
}
