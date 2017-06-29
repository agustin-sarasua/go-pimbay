package db

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/agustin-sarasua/pimbay/model"
)

// Create a User account uid = user datastore id
func (db *datastoreDB) SaveAccount(b *model.Account, uid int64) (id int64, err error) {
	ctx := context.Background()
	uk := datastore.IDKey("User", uid, nil)
	k := datastore.IncompleteKey("Account", uk)
	k, err = db.client.Put(ctx, k, b)
	if err != nil {
		return 0, fmt.Errorf("datastoredb: could not put Account: %v", err)
	}
	fmt.Println(k)
	return k.ID, nil
}

func (db *datastoreDB) ListUserAccounts(uid int64) (as []*model.Account, err error) {
	ctx := context.Background()
	uk := datastore.IDKey("User", uid, nil)
	q := datastore.NewQuery("Account").Ancestor(uk)
	var accounts []*model.Account
	_, err = db.client.GetAll(ctx, q, &accounts)
	return accounts, err
}

func (db *datastoreDB) GetAccount(id int64) (*model.Account, error) {
	ctx := context.Background()
	k := datastore.IDKey("Account", id, nil)
	u := &model.Account{}
	if err := db.client.Get(ctx, k, u); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get Account: %v", err)
	}
	u.ID = id
	return u, nil
}
