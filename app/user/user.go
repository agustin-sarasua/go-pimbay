package user

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
