package model

import (
	"time"
)

type User struct {
	ID          int64     `json:"id" bson:"_id"`
	FirebaseID  string    `json:"firebaseID" bson:"firebaseID"`
	Name        string    `json:"name" bson:"name"`
	LastName    string    `json:"lastName" bson:"lastName"`
	Email       string    `json:"email" bson:"email"`
	Birthdate   time.Time `json:"birthdate,omitempty" bson:"birthdate,omitempty"`
	CreatedDate time.Time `json:"createdDate" bson:"createdDate"`
	Sex         string    `json:"sex,omitempty" bson:"sex,omitempty"`
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

type UserDatabase interface {
	SaveUser(u *User) (id int64, e error)
	GetUser(id int64) (*User, error)
	Close()
}
