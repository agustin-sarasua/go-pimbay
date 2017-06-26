package model

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

const (
	accCollection = "account"
)

// Ensure mongoDB conforms to the UserDatabase interface.
var _ AccountDatabase = &mongoDB{}

func (db *mongoDB) SaveAccount(a *Account) (id string, e error) {
	session := db.conn.Copy()
	defer session.Close()

	c := session.DB("gopimbay").C(accCollection)
	err := c.Insert(a)
	if err != nil {
		if mgo.IsDup(err) {
			panic(err)
		}
		log.Println("Failed insert User: ", err)
	}
	return a.ID, nil
}
