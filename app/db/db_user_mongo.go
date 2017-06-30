package db

import (
	"log"

	"github.com/agustin-sarasua/pimbay/app/model"

	mgo "gopkg.in/mgo.v2"
)

const (
	userCollection = "user"
)

// Ensure mongoDB conforms to the UserDatabase interface.
//var _ UserDatabase = &mongoDB{}

func (db *mongoDB) SaveUser(u *model.User) (id int64, e error) {
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

func (db *mongoDB) GetUser(id int64) (*model.User, error) {
	return nil, nil
}

// Close closes the database.
func (db *mongoDB) Close() {
	db.conn.Close()
}
