package reports

import (
	"fmt"

	"golang.org/x/net/context"

	"cloud.google.com/go/datastore"
)

var (
	ReportDB ReportDatabase
	_        ReportDatabase = &datastoreDB{}
)

type ReportDatabase interface {
	Cleanup()
}

/**
	Datastore
**/

type datastoreDB struct {
	client *datastore.Client
}

// Ensure datastoreDB conforms to the UserDatabase  and AccountDatabase interface.
var _ ReportDatabase = &datastoreDB{}

// See the datastore and google packages for details on creating a suitable Client:
// https://godoc.org/cloud.google.com/go/datastore
func NewDatastoreDB(ctx context.Context, client *datastore.Client) (ReportDatabase, error) {
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
}
