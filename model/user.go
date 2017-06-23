package model

import (
	"time"
)

type User struct {
	ID          string    `json:"id" bson:"_id"`
	Name        string    `json:"name" bson:"name"`
	LastName    string    `json:"lastName" bson:"lastName"`
	Email       string    `json:"email" bson:"email"`
	Birthdate   time.Time `json:"birthdate,omitempty" bson:"birthdate,omitempty"`
	CreatedDate time.Time `json:"createdDate" bson:"createdDate"`
	Sex         string    `json:"sex,omitempty" bson:"sex,omitempty"`
}

type Account struct {
	ID          string    `json:"id,omitempty" bson:"_id"`
	Name        string    `json:"name" bson:"name"`
	CreatedDate time.Time `json:"createdDate,omitempty" bson:"createdDate"`
}

type CreditCard struct {
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
	CountryCode string `json:"countryCode" bson:"countryCode"`
}

/*

Number struct {
		Length int `json:"length"`
		Luhn bool `json:"luhn"`
	} `json:"number"`
	Scheme string `json:"scheme"`
	Type string `json:"type"`
	Brand string `json:"brand"`
	Prepaid bool `json:"prepaid"`
	Country struct {
		Numeric string `json:"numeric"`
		Alpha2 string `json:"alpha2"`
		Name string `json:"name"`
		Emoji string `json:"emoji"`
		Currency string `json:"currency"`
		Latitude int `json:"latitude"`
		Longitude int `json:"longitude"`
	} `json:"country"`


*/

/*
type Account struct {
	ID          string
	Name        string
	CreatedDate util.DateTime
	Country     string
	CreditCards *[]CreditCard
	Expenses    *[]Expense
}

type CreditCard struct {
	ID         string
	LastDigits string
	Bin        string
	Bank       string
}

type Expense struct {
	ID          string
	Amount      util.BigDecimal
	Currency    string
	DateTime    util.DateTime
	Report      *Report
	IndexedDate util.DateTime
	Category    string
	Business    *Business
}

type Business struct {
	ID         string
	BusinessID string
	Name       string
	Category   string
	Address    *Address
}

type Address struct {
}
*/
