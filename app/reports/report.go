package reports

import (
	"time"

	"github.com/shopspring/decimal"
)

type Report interface {
	GetTransactionLines() []string
	GetLines() []string
	ParseLineDetail(l string) (time.Time, string, string, string, decimal.Decimal)

	GetReportLines() []*ReportLine
	IsTransactionLine(l string) bool
	GetTransactionDate(l string) time.Time
}

type ReportLine struct {
	Date         time.Time
	LastDigits   string
	BusinessName string
	Currency     string
	Amount       decimal.Decimal
}
