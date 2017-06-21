package model

import (
	"fmt"
	"regexp"
	"strings"
)

type Report interface {
	GetTransactionLines() []string
	GetLines() []string
}

type ReportLineValidator interface {
	IsTransactionLine() bool
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
		if r.IsTransactionLine(strings.TrimSpace(l)) {
			rs = append(rs, l)
		} else {
			fmt.Print("NO TX LINE: ")
			fmt.Println(l)
		}
	}
	return rs
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
