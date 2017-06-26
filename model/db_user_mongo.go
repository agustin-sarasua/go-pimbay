package model

import (
	"log"
	mgo "gopkg.in/mgo.v2"
)

const (
	userCollection = "user"
)

// Ensure mongoDB conforms to the UserDatabase interface.
var _ UserDatabase = &mongoDB{}


func (db *mongoDB) SaveUser(u *User) (id string, e error) {
	session := db.conn.Copy()
	defer session.Close()

	c := session.DB("gopimbay").C(userCollection)
	err := c.Insert(u)
	if err != nil {
		if mgo.IsDup(err) {
			panic(err)
		}
		log.Println("Failed insert User: ", err)
	}
	return u.ID, nil
}
