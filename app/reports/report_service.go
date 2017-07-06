package reports

import (
	"bytes"
	"fmt"

	"github.com/ledongthuc/pdf"
)

func ReadPdf(path string) (string, error) {
	r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	var textBuilder bytes.Buffer
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		textBuilder.WriteString(p.GetPlainText("\n"))
	}
	return textBuilder.String(), nil
}

func PrintValidReportLines(r Report) {
	ls := r.GetTransactionLines()
	for _, l := range ls {
		t, ld, bn, c, a := r.ParseLineDetail(l)
		fmt.Println(t, " ", ld, " ", bn, " ", c, a)
	}
}

func ProcessReport(r *Report) {

}
