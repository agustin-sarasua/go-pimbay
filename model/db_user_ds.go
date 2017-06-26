package model

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
)

type datastoreDB struct {
	client *datastore.Client
}

var _ UserDatabase = &datastoreDB{}

// See the datastore and google packages for details on creating a suitable Client:
// https://godoc.org/cloud.google.com/go/datastore
func NewDatastoreDB(client *datastore.Client) (UserDatabase, error) {
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

func (db *datastoreDB) SaveUser(b *User) (id string, err error) {
	ctx := context.Background()
	k := datastore.IncompleteKey("User", nil)
	k, err = db.client.Put(ctx, k, b)
	if err != nil {
		return "", fmt.Errorf("datastoredb: could not put USer: %v", err)
	}
	return k.String(), nil
}
