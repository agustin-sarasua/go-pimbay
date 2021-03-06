package account

import (
	"time"
)

type Account struct {
	ID          int64     `json:"id,omitempty" bson:"_id"`
	Name        string    `json:"name" bson:"name"`
	CreatedDate time.Time `json:"createdDate,omitempty" bson:"createdDate"`
}

type CreditCard struct {
	ID         string `json:"id" bson:"_id"`
	Bin        string `json:"bin" bson:"bin"`
	LastDigits string `json:"lastDigits" bson:"lastDigits"`
	Type       string `json:"type" bson:"type"`
	Prepaid    bool   `json:"prepaid" bson:"prepaid"`
	Brand      string `json:"brand" bson:"brand"`
	Bank       struct {
		Name  string `json:"name"`
		URL   string `json:"url"`
		Phone string `json:"phone"`
		City  string `json:"city"`
	} `json:"bank"`
	CountryCode string    `json:"countryCode" bson:"countryCode"`
	CreatedDate time.Time `json:"createdDate" bson:"createdDate"`
}
