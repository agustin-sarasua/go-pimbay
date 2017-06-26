package model

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

const (
	dbName = "gopimbay"
)

type mongoDB struct {
	conn *mgo.Session
}

// Ensure mongoDB conforms to the UserDatabase interface.
var _ UserDatabase = &mongoDB{}

func NewMongoDB(addr, db, username, pwd, c string) (UserDatabase, error) {
	// Mongo
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{addr},
		Timeout:  60 * time.Second,
		Database: db,
		Username: username,
		Password: pwd,
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	return &mongoDB{conn: session}, nil
}
