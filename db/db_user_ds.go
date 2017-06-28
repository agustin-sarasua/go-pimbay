package db

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/agustin-sarasua/pimbay/model"
)

type datastoreDB struct {
	client *datastore.Client
}

var _ UserDatabase = &datastoreDB{}

// See the datastore and google packages for details on creating a suitable Client:
// https://godoc.org/cloud.google.com/go/datastore
func NewDatastoreDB(client *datastore.Client) (model.UserDatabase, error) {
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

func (db *datastoreDB) SaveUser(b *model.User) (id int64, err error) {
	ctx := context.Background()
	k := datastore.IncompleteKey("User", nil)
	k, err = db.client.Put(ctx, k, b)
	if err != nil {
		return 0, fmt.Errorf("datastoredb: could not put User: %v", err)
	}
	fmt.Println(k)
	return k.ID, nil
}

func (db *datastoreDB) GetUser(id int64) (*model.User, error) {
	ctx := context.Background()
	k := datastore.IDKey("User", id, nil)
	u := &model.User{}
	if err := db.client.Get(ctx, k, u); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get User: %v", err)
	}
	u.ID = id
	return u, nil
}

// Close closes the database.
func (db *datastoreDB) Close() {
	// No op.
}
