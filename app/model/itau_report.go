package model

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type ItauReport struct {
	Content     string
	Lines       []string
	ReportLines []*ItauReportLine
}

type ItauReportLine struct {
	Date         time.Time
	LastDigits   string
	BusinessName string
	Currency     string
	Amount       decimal.Decimal
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

func (r ItauReport) ParseLineDetail(l string) (time.Time, string, string, string, decimal.Decimal) {
	c, d := r.GetTransactionAmount(l)
	return r.GetTransactionDate(l), r.GetCardLastDigits(l), r.GetBusinessName(l), c, d
}

func (r ItauReport) GetTransactionDate(l string) time.Time {
	var lineSearch = regexp.MustCompile(`^([0-3][0-9])(\ )([0-1][0-9])(\ )([0-2][0-9])`)
	layout := "02 01 06"
	t := lineSearch.FindString(l)
	rs, _ := time.Parse(layout, t)
	return rs
}

func (r ItauReport) GetCardLastDigits(l string) string {
	var lineDateSearch = regexp.MustCompile(`^([0-3][0-9])(\ )([0-1][0-9])(\ )([0-2][0-9])`)
	var lastDigitsSearch = regexp.MustCompile(`^\d{4}`)
	t := strings.TrimSpace(lineDateSearch.ReplaceAllString(l, ""))
	lastDigits := t[0:4]
	if lastDigitsSearch.FindStringIndex(lastDigits) != nil {
		return lastDigits
	}
	return "NOT_FOUND"
}

func (r ItauReport) GetBusinessName(l string) string {
	var lineSearch = regexp.MustCompile(`^([0-3][0-9])(\ )([0-1][0-9])(\ )([0-2][0-9])(\ ){2}\d{4}`)
	t := strings.TrimSpace(lineSearch.ReplaceAllString(l, ""))
	a := strings.Split(t, "     ")
	return strings.TrimSpace(a[0])
}

func (r ItauReport) GetTransactionAmount(l string) (string, decimal.Decimal) {
	var lineSearch = regexp.MustCompile(`^([0-3][0-9])(\ )([0-1][0-9])(\ )([0-2][0-9])(\ ){2}\d{4}`)
	var amountSearch = regexp.MustCompile(`\d*(\,)\d{2}$`)
	t := strings.TrimSpace(lineSearch.ReplaceAllString(l, ""))
	m := amountSearch.FindAllString(t, -1)
	if m == nil {
		panic("No amount found")
	}
	d, e := decimal.NewFromString(strings.Replace(m[0], ",", ".", 1))
	if e != nil {
		panic(e)
	}
	var c string
	if len(t) == 93 {
		c = "USD"
	} else if len(t) == 79 {
		c = "UYP"
	} else {
		fmt.Println("Largo extranio")
	}
	return c, d
}
