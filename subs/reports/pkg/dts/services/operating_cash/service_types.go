package operating_cash

import "time"

type Response struct {
	Data []Data `json:"data"`
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

type Data struct {
	RecordDate                        time.Time `json:"record_date"`
	AccountType                       string    `json:"account_type"`
	TransactionType                   string    `json:"transaction_type"`
	TransactionCategory               string    `json:"transaction_catg"`
	TransactionCategoryDescription    string    `json:"transaction_catg_desc"`
	TransactionTodayAmount            int       `json:"transaction_today_amt"`
	TransactionMonthToDateAmount      int       `json:"transaction_mtd_amt"`
	TransactionFiscalYearToDateAmount int       `json:"transaction_fytd_amt"`
	TableNumber                       string    `json:"table_nbr"`
	TableName                         string    `json:"table_nm"`
	SourceLineNumber                  string    `json:"src_line_nbr"`
	RecordFiscalYear                  int       `json:"record_fiscal_year"`
	RecordFiscalQuarter               int       `json:"record_fiscal_quarter"`
	RecordCalendarYear                int       `json:"record_calendar_year"`
	RecordCalendarQuarter             int       `json:"record_calendar_quarter"`
	RecordCalendarMonth               int       `json:"record_calendar_month"`
	RecordCalendarDay                 int       `json:"record_calendar_day"`
}

func into(result response) Response {
	out := Response{
		Meta:  result.Meta,
		Links: result.Links,
	}

	data := make([]Data, len(result.Data))

	for i, d := range result.Data {
		data[i] = Data{
			RecordDate:                        d.RecordDate.Into(),
			AccountType:                       d.AccountType,
			TransactionType:                   d.TransactionType,
			TransactionCategory:               d.TransactionCategory,
			TransactionCategoryDescription:    d.TransactionCategoryDescription,
			TransactionTodayAmount:            d.TransactionTodayAmount.Into(),
			TransactionMonthToDateAmount:      d.TransactionMonthToDateAmount.Into(),
			TransactionFiscalYearToDateAmount: d.TransactionFiscalYearToDateAmount.Into(),
			TableNumber:                       d.TableNumber,
			TableName:                         d.TableName,
			SourceLineNumber:                  d.SourceLineNumber,
			RecordFiscalYear:                  d.RecordFiscalYear.Into(),
			RecordFiscalQuarter:               d.RecordFiscalQuarter.Into(),
			RecordCalendarYear:                d.RecordCalendarYear.Into(),
			RecordCalendarQuarter:             d.RecordCalendarQuarter.Into(),
			RecordCalendarMonth:               d.RecordCalendarMonth.Into(),
			RecordCalendarDay:                 d.RecordCalendarDay.Into(),
		}
	}

	out.Data = data

	return out
}
