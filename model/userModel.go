package model

import (
	"time"

	"github.com/agustin-sarasua/pimbay/util"
)

type User struct {
	ID          string
	Name        string
	LastName    string
	Birthdate   time.Time
	CreatedDate time.Time
	Sex         string
}

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

type Report struct {
	ID          string
	Name        string
	Account     *Account
	CreditCard  *CreditCard
	DateTime    util.DateTime
	IndexedDate util.DateTime
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
