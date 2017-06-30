package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Report interface {
	GetTransactionLines() []string
	GetLines() []string
	ParseLineDetail(l string) (time.Time, string, string, string, decimal.Decimal)
}

type ReportLine interface {
	IsTransactionLine() bool
	GetTransactionDate(l string) time.Time
}
