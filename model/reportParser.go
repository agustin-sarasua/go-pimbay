package model

import (
	"regexp"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type Report interface {
	GetTransactionLines() []string
	GetLines() []string
	ParseLineDetail(l string) (time.Time, string, string, decimal.Decimal)
}

type ReportLine interface {
	IsTransactionLine() bool
	GetTransactionDate(l string) time.Time
}

type ItauReport struct {
	Content string
	Lines   []string
}

func (r ItauReport) IsTransactionLine(l string) bool {
	var lineSearch = regexp.MustCompile(`^([0-3][0-9])(\ )([0-1][0-9])(\ )([0-2][0-9])(\ ){2}\d{4}`)
	return lineSearch.FindStringIndex(l) != nil && lineSearch.FindStringIndex(l)[0] == 0
}

func (r ItauReport) GetLines() []string {
	return r.Lines
}

func (r ItauReport) GetTransactionLines() []string {
	rs := r.Lines[:0]
	for _, l := range r.Lines {
		tl := strings.TrimSpace(l)
		if r.IsTransactionLine(tl) {
			rs = append(rs, tl)
		}
	}
	return rs
}

func (r ItauReport) ParseLineDetail(l string) (time.Time, string, string, decimal.Decimal) {
	return r.GetTransactionDate(l), r.GetCardLastDigits(l), "detail", decimal.Decimal{}
}

func (r ItauReport) GetTransactionDate(l string) time.Time {
	var lineSearch = regexp.MustCompile(`^([0-3][0-9])(\ )([0-1][0-9])(\ )([0-2][0-9])(\ ){2}\d{4}`)
	layout := "02 01 06"
	t := lineSearch.FindString(l)
	rs, _ := time.Parse(layout, t)
	return rs
}

func (r ItauReport) GetCardLastDigits(l string) string {
	var lineDateSearch = regexp.MustCompile(`^([0-3][0-9])(\ )([0-1][0-9])(\ )([0-2][0-9])(\ ){2}\d{4}`)
	var lastDigitsSearch = regexp.MustCompile(`^\d{4}`)
	t := strings.TrimSpace(lineDateSearch.ReplaceAllString(l, ""))
	lastDigits := t[0:3]
	if lastDigitsSearch.FindStringIndex(lastDigits) != nil {
		return lastDigits
	}
	return "NOT_FOUND"
}

/* type ReportParser interface {
	FilterReportLines() []string
}

type ReportLineParser interface {
	IsValidTxLine() bool
}

type ItauReport struct {
	Content string
	Lines   []ItauReportLine
}

type ItauReportLine struct{ string }

func (l ItauReportLine) IsValidTxLine() bool {
	var lineSearch = regexp.MustCompile(`^([0-3][0-9])(\ )([0-1][0-9])(\ )([0-2][0-9])(\ ){2}\d{4}`)
	return lineSearch.FindStringIndex(l.string) != nil && lineSearch.FindStringIndex(l.string)[0] == 0
}

func (r ItauReport) FilterReportLines() []string {
	rs := r.Lines[:0]
	for _, l := range r.Lines {
		if IsValidTxLine(l) {
			rs = append(rs, l)
		}
	}
	return rs
}

func
*/
