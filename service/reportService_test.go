package service_test

import (
	"fmt"
	"strings"
	"testing"

	"regexp"

	"github.com/agustin-sarasua/pimbay/model"
	"github.com/agustin-sarasua/pimbay/service"
)

const (
	usdLineLength = 111
	uyLineLenght  = 97
)

func TestReadPdf(t *testing.T) {
	var lineSearch = regexp.MustCompile(`^([0-3][0-9])(\ )([0-1][0-9])(\ )([0-2][0-9])(\ ){2}\d{4}`)
	fmt.Println("Running test")
	content, err := service.ReadPdf("../resources/Est_Cta_Visa_201701.pdf")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			fmt.Println("Break!")
		} else {
			if lineSearch.FindStringIndex(line) != nil && lineSearch.FindStringIndex(line)[0] == 0 {
				fmt.Print(lineSearch.FindStringIndex(line))
				fmt.Print(len(line))
				fmt.Println(lineSearch.FindAllString(line, -1))
			}
		}

		//regexp.Match(validID, line)
		//for i := 0; i < len(line); i++ {
		//fmt.Printf("%x ", line[i])

		//}
		fmt.Println("--")
	}
	//fmt.Println(content)
	return
}

func TestReportParser(t *testing.T) {
	content, err := service.ReadPdf("../resources/Est_Cta_Visa_201701.pdf")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\n")

	report := model.ItauReport{Content: content, Lines: lines}
	service.PrintValidReportLines(report)
}
