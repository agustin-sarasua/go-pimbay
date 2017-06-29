package db

import (
	"fmt"

	"golang.org/x/net/context"

	"cloud.google.com/go/datastore"
	"github.com/agustin-sarasua/pimbay/model"
)

// Create a User account uid = user datastore id
func (db *datastoreDB) SaveAccount(ctx context.Context, b *model.Account, uid int64) (id int64, err error) {
	uk := datastore.IDKey("User", uid, nil)
	k := datastore.IncompleteKey("Account", uk)
	k, err = db.client.Put(ctx, k, b)
	if err != nil {
		return 0, fmt.Errorf("datastoredb: could not put Account: %v", err)
	}
	fmt.Println(k)
	return k.ID, nil
}

func (db *datastoreDB) ListUserAccounts(ctx context.Context, uid int64) (as []*model.Account, err error) {
	uk := datastore.IDKey("User", uid, nil)
	q := datastore.NewQuery("Account").Ancestor(uk)
	var accounts []*model.Account
	_, err = db.client.GetAll(ctx, q, &accounts)
	return accounts, err
}

func (db *datastoreDB) GetAccount(ctx context.Context, id int64) (*model.Account, error) {
	k := datastore.IDKey("Account", id, nil)
	u := &model.Account{}
	if err := db.client.Get(ctx, k, u); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get Account: %v", err)
	}
	u.ID = id
	return u, nil
}
