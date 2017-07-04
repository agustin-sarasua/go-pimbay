package db

import (
	"fmt"

	"cloud.google.com/go/datastore"

	"golang.org/x/net/context"

	mgo "gopkg.in/mgo.v2"

	//"cloud.google.com/go/datastore"

	"github.com/agustin-sarasua/pimbay/app/model"
)

type Database interface {
	SaveUser(ctx context.Context, u *model.User) (id int64, e error)
	GetUser(ctx context.Context, id int64) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByFirebaseID(ctx context.Context, fID string) (*model.User, error)
	DeleteUser(id int64) error
	Close()

	SaveAccount(ctx context.Context, a *model.Account, uid int64) (id int64, e error)
	GetAccount(ctx context.Context, id int64) (*model.Account, error)
	ListUserAccounts(ctx context.Context, uid int64) (as []*model.Account, err error)

	Cleanup()
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
func NewDatastoreDB(ctx context.Context, client *datastore.Client) (Database, error) {
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

func (db *datastoreDB) Cleanup() {
	fmt.Println("Cleaning up datastore...")
	ctx := context.Background()
	q := datastore.NewQuery("User")
	var usersData []*model.User
	ks, _ := db.client.GetAll(ctx, q, &usersData)
	fmt.Printf("Keys to delete %d\n", len(ks))
	db.client.DeleteMulti(ctx, ks)
	fmt.Println("Datastore cleaned up")
}
