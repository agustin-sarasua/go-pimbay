package reports

import "time"

type SaveReportRequest struct {
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}
