package reports_test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"cloud.google.com/go/storage"

	"github.com/agustin-sarasua/pimbay/app/reports"

	"regexp"
)

const (
	usdLineLength = 111
	uyLineLenght  = 97
)

func TestReadPdf(t *testing.T) {
	var lineSearch = regexp.MustCompile(`^([0-3][0-9])(\ )([0-1][0-9])(\ )([0-2][0-9])(\ ){2}\d{4}`)
	fmt.Println("Running test")
	content, err := reports.ReadPdf("../resources/Est_Cta_Visa_201701.pdf")
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
		fmt.Println("--")
	}
	return
}

func TestReportParser(t *testing.T) {
	content, err := reports.ReadPdf("../resources/Est_Cta_Visa_201701.pdf")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\n")

	report := reports.ItauReport{Content: content, Lines: lines}
	reports.PrintValidReportLines(report)
}

func TestFormatShortYear(t *testing.T) {
	value := "22 09 17"
	layout := "02 01 06"
	r, _ := time.Parse(layout, value)
	fmt.Println(r)

}

func TestReadFile(t *testing.T) {
	reports.StorageBucketName = "pimbay-accounting.appspot.com"
	reports.StorageBucket, _ = configureStorage(reports.StorageBucketName)

	s, _ := reports.ReadFile("f0e9d430-496f-4858-92d0-e8140c2f3984.pdf", context.Background())
	fmt.Println(s)
}

func configureStorage(bucketID string) (*storage.BucketHandle, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.Bucket(bucketID), nil
}
