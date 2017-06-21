package service

import (
	"bytes"

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

/*func ReadPdf2(path string) (string, error) {
	r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		var lastTextStyle pdf.Text
		texts := p.Content().Text
		for _, text := range texts {
			if isSameSentence(text, lastTextStyle) {
				lastTextStyle.S = lastTextStyle.S + text.S
			} else {
				fmt.Printf("Font: %s, Font-size: %f, x: %f, y: %f, content: %s \n", lastTextStyle.Font, lastTextStyle.FontSize, lastTextStyle.X, lastTextStyle.Y, lastTextStyle.S)
				lastTextStyle = text
			}
		}
	}
	return "", nil
}*/
