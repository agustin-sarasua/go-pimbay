package db

import (
	"context"
	"fmt"

	mgo "gopkg.in/mgo.v2"

	"cloud.google.com/go/datastore"
	"github.com/agustin-sarasua/pimbay/model"
)

type Database interface {
	SaveUser(u *model.User) (id int64, e error)
	GetUser(id int64) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	Close()

	SaveAccount(a *model.Account, uid int64) (id int64, e error)
	GetAccount(id int64) (*model.Account, error)
	ListUserAccounts(uid int64) (as []*model.Account, err error)
}

/**
	MONGO DB
**/

const (
	dbName = "gopimbay"
)

type mongoDB struct {
	conn *mgo.Session
}

// Ensure mongoDB conforms to the UserDatabase interface.
//var _ Database = &mongoDB{}

// func NewMongoDB(addr, db, username, pwd, c string) (Database, error) {
// 	// Mongo
// 	mongoDBDialInfo := &mgo.DialInfo{
// 		Addrs:    []string{addr},
// 		Timeout:  60 * time.Second,
// 		Database: db,
// 		Username: username,
// 		Password: pwd,
// 	}
// 	session, err := mgo.DialWithInfo(mongoDBDialInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	//defer session.Close()
// 	session.SetMode(mgo.Monotonic, true)

// 	return &mongoDB{conn: session}, nil
// }

/**
	Datastore
**/

type datastoreDB struct {
	client *datastore.Client
}

// Ensure datastoreDB conforms to the UserDatabase  and AccountDatabase interface.
var _ Database = &datastoreDB{}

// See the datastore and google packages for details on creating a suitable Client:
// https://godoc.org/cloud.google.com/go/datastore
func NewDatastoreDB(client *datastore.Client) (Database, error) {
	ctx := context.Background()
	// Verify that we can communicate and authenticate with the datastore service.
	t, err := client.NewTransaction(ctx)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
	}
	if err := t.Rollback(); err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
	}
	return &datastoreDB{
		client: client,
	}, nil
}
