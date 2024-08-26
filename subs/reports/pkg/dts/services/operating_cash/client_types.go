package operating_cash

import (
	"fmt"
	"strings"
	"time"

	"github.com/garrettladley/left-arrow/subs/reports/pkg/dts/constants"
	"github.com/garrettladley/left-arrow/subs/reports/pkg/dts/types"
)

type response struct {
	Data []struct {
		RecordDate                        types.Date      `json:"record_date"`
		AccountType                       string          `json:"account_type"`
		TransactionType                   string          `json:"transaction_type"`
		TransactionCategory               string          `json:"transaction_catg"`
		TransactionCategoryDescription    string          `json:"transaction_catg_desc"`
		TransactionTodayAmount            types.StringInt `json:"transaction_today_amt"`
		TransactionMonthToDateAmount      types.StringInt `json:"transaction_mtd_amt"`
		TransactionFiscalYearToDateAmount types.StringInt `json:"transaction_fytd_amt"`
		TableNumber                       string          `json:"table_nbr"`
		TableName                         string          `json:"table_nm"`
		SourceLineNumber                  string          `json:"src_line_nbr"`
		RecordFiscalYear                  types.StringInt `json:"record_fiscal_year"`
		RecordFiscalQuarter               types.StringInt `json:"record_fiscal_quarter"`
		RecordCalendarYear                types.StringInt `json:"record_calendar_year"`
		RecordCalendarQuarter             types.StringInt `json:"record_calendar_quarter"`
		RecordCalendarMonth               types.StringInt `json:"record_calendar_month"`
		RecordCalendarDay                 types.StringInt `json:"record_calendar_day"`
	} `json:"data"`
	Meta struct {
		Count      int `json:"count"`
		TotalCount int `json:"total-count"`
		TotalPages int `json:"total-pages"`
	} `json:"meta"`
	Links struct {
		Self  string `json:"self"`
		First string `json:"first"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
		Last  string `json:"last"`
	} `json:"links"`
}

type params struct {
	Fields     []string
	Filters    []filter
	Sorts      []sort // Sorts are applied in the order they are specified
	PageSize   int
	PageNumber int
}

type filterModifier string

const (
	filterModifierLessThan           filterModifier = "lt"
	filterModifierLessThanOrEqual    filterModifier = "lte"
	filterModifierGreaterThan        filterModifier = "gt"
	filterModifierGreaterThanOrEqual filterModifier = "gte"
	filterModifierEqual              filterModifier = "eq"
	filterModifierIn                 filterModifier = "in"
)

type filter struct {
	Field    string
	Modifier filterModifier
	Value    string
}

func (f *filter) string() string {
	return fmt.Sprintf("%s:%s:%s", f.Field, f.Modifier, f.Value)
}

func before(value time.Time) filter {
	return filter{
		Field:    "record_date",
		Modifier: filterModifierLessThanOrEqual,
		Value:    value.Format(constants.YYYY_MM_DD),
	}
}

func after(value time.Time) filter {
	return filter{
		Field:    "record_date",
		Modifier: filterModifierGreaterThanOrEqual,
		Value:    value.Format(constants.YYYY_MM_DD),
	}
}

func in(values ...string) filter {
	return filter{
		Field:    "record_date",
		Modifier: filterModifierIn,
		Value:    strings.Join(values, ","),
	}
}

func xfields(values ...string) filter {
	return filter{
		Field:    "fields",
		Modifier: filterModifierEqual,
		Value:    strings.Join(values, ","),
	}
}

type sort struct {
	IsPositive bool
	Field      string
}

func (s *sort) string() string {
	if s.IsPositive {
		return s.Field
	}
	return fmt.Sprintf("-%s", s.Field)
}
