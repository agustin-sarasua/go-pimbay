package account

import (
	"fmt"

	"golang.org/x/net/context"

	"cloud.google.com/go/datastore"
)

var (
	AccountDB AccountDatabase
	_         AccountDatabase = &datastoreDB{}
)

type AccountDatabase interface {
	SaveAccount(ctx context.Context, a *Account, uid int64) (id int64, e error)
	GetAccount(ctx context.Context, id int64) (*Account, error)
	ListUserAccounts(ctx context.Context, uid int64) (as []*Account, err error)

	Cleanup()
}

// Create a User account uid = user datastore id
func (db *datastoreDB) SaveAccount(ctx context.Context, b *Account, uid int64) (id int64, err error) {
	uk := datastore.IDKey("User", uid, nil)
	k := datastore.IncompleteKey("Account", uk)
	k, err = db.client.Put(ctx, k, b)
	if err != nil {
		return 0, fmt.Errorf("datastoredb: could not put Account: %v", err)
	}
	fmt.Println(k)
	return k.ID, nil
}

func (db *datastoreDB) ListUserAccounts(ctx context.Context, uid int64) (as []*Account, err error) {
	uk := datastore.IDKey("User", uid, nil)
	q := datastore.NewQuery("Account").Ancestor(uk)
	var accounts []*Account
	_, err = db.client.GetAll(ctx, q, &accounts)
	return accounts, err
}

func (db *datastoreDB) GetAccount(ctx context.Context, id int64) (*Account, error) {
	k := datastore.IDKey("Account", id, nil)
	u := &Account{}
	if err := db.client.Get(ctx, k, u); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get Account: %v", err)
	}
	u.ID = id
	return u, nil
}

/**
	Datastore
**/

type datastoreDB struct {
	client *datastore.Client
}

// Ensure datastoreDB conforms to the UserDatabase  and AccountDatabase interface.
var _ AccountDatabase = &datastoreDB{}

// See the datastore and google packages for details on creating a suitable Client:
// https://godoc.org/cloud.google.com/go/datastore
func NewDatastoreDB(ctx context.Context, client *datastore.Client) (AccountDatabase, error) {
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
