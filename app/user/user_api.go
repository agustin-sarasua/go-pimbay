package user

import "time"

type SignupUserRestMsg struct {
	Name      string    `json:"name" bson:"name"`
	LastName  string    `json:"lastName" bson:"lastName"`
	Email     string    `json:"email" bson:"email"`
	Birthdate time.Time `json:"birthdate,omitempty" bson:"birthdate,omitempty"`
	Sex       string    `json:"sex,omitempty" bson:"sex,omitempty"`
	Password  string    `json:"password" bson:"password"`
}
