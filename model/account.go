package model

import "time"

type Account struct {
	ID          string        `json:"id,omitempty" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	CreatedDate time.Time     `json:"createdDate,omitempty" bson:"createdDate"`
	CreditCards []*CreditCard `json:"creditCards,omitempty" bson:"creditCards"`
}

type AccountDatabase interface {
	SaveAccount(a *Account) (id string, e error)
}
