package model

import "time"

type Account struct {
	ID          int64         `json:"id,omitempty" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	CreatedDate time.Time     `json:"createdDate,omitempty" bson:"createdDate"`
	CreditCards []*CreditCard `json:"creditCards,omitempty" bson:"creditCards"`
}
