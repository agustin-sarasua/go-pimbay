package service_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/agustin-sarasua/pimbay/service"
)

func TestReadPdf(t *testing.T) {
	//var validID = regexp.MustCompile(`^(([1-9])|([0][1-9])|([1-2][0-9])|([3][0-1]))\ (([1-9])|([0][1-9])|([1-2][0-9])|([3][0-1]))\ (([1-9])|([0][1-9])|([1-2][0-9])|([3][0-1]))(\ *)\d{4}.*$
	//`)
	fmt.Println("Running test")
	content, err := service.ReadPdf("../resources/Est_Cta_Visa_201701.pdf")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		line = strings.TrimLeft(line, " ")

		fmt.Println(line)
		//regexp.Match(validID, line)
		for i := 0; i < len(line); i++ {
			//fmt.Printf("%x ", line[i])

		}
		fmt.Println("--")
	}
	//fmt.Println(content)
	return
}
