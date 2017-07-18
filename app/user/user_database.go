package user

import (
	"fmt"

	"golang.org/x/net/context"

	"cloud.google.com/go/datastore"
)

var (
	UserDB UserDatabase
	_      UserDatabase = &datastoreDB{}
)

type UserDatabase interface {
	SaveUser(ctx context.Context, u *User) (id int64, e error)
	GetUser(ctx context.Context, id int64) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByFirebaseID(ctx context.Context, fID string) (*User, error)
	DeleteUser(id int64) error
	Close()

	Cleanup()
}

type datastoreDB struct {
	client *datastore.Client
}

func (db *datastoreDB) SaveUser(ctx context.Context, b *User) (id int64, err error) {
	k := datastore.IncompleteKey("User", nil)
	k, err = db.client.Put(ctx, k, b)
	if err != nil {
		return 0, fmt.Errorf("datastoredb: could not put User: %v", err)
	}
	fmt.Println(k)
	return k.ID, nil
}

func (db *datastoreDB) GetUser(ctx context.Context, id int64) (*User, error) {
	k := datastore.IDKey("User", id, nil)
	u := &User{}
	if err := db.client.Get(ctx, k, u); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get User: %v", err)
	}
	u.ID = id
	return u, nil
}

func (db *datastoreDB) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	q := datastore.NewQuery("User").Filter("Email =", email)
	var us []*User
	ks, _ := db.client.GetAll(ctx, q, &us)
	if ks != nil && len(ks) > 0 {
		us[0].ID = ks[0].ID
		return us[0], nil
	}
	return nil, nil
}

func (db *datastoreDB) GetUserByFirebaseID(ctx context.Context, fID string) (*User, error) {
	q := datastore.NewQuery("User").Filter("FirebaseID =", fID)
	var us []*User
	ks, _ := db.client.GetAll(ctx, q, &us)
	if ks != nil && len(ks) > 0 {
		us[0].ID = ks[0].ID
		return us[0], nil
	}
	return nil, nil
}

func (db *datastoreDB) DeleteUser(id int64) error {
	ctx := context.Background()
	return db.client.Delete(ctx, datastore.IDKey("User", id, nil))
}

// Close closes the database.
func (db *datastoreDB) Close() {
	// No op.

}

func (db *datastoreDB) Cleanup() {
	fmt.Println("Cleaning up datastore...")
	ctx := context.Background()
	q := datastore.NewQuery("User")
	var usersData []*User
	ks, _ := db.client.GetAll(ctx, q, &usersData)
	fmt.Printf("Keys to delete %d\n", len(ks))
	db.client.DeleteMulti(ctx, ks)
	fmt.Println("Datastore cleaned up")
}

// See the datastore and google packages for details on creating a suitable Client:
// https://godoc.org/cloud.google.com/go/datastore
func NewDatastoreDB(ctx context.Context, client *datastore.Client) (UserDatabase, error) {
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
