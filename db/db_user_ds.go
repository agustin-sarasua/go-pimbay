package db

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/agustin-sarasua/pimbay/model"
)

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

func (db *datastoreDB) GetUserByEmail(email string) (*model.User, error) {
	ctx := context.Background()
	q := datastore.NewQuery("User").Filter("Email =", email)
	var us []*model.User
	db.client.GetAll(ctx, q, &us)
	return us[0], nil
}

// Close closes the database.
func (db *datastoreDB) Close() {
	// No op.

}
