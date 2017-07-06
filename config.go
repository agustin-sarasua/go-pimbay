package pimbay

import (
	"context"
	"flag"
	"fmt"
	"os"

	"cloud.google.com/go/datastore"
	"cloud.google.com/go/storage"
	"github.com/agustin-sarasua/pimbay/app/account"
	"github.com/agustin-sarasua/pimbay/app/api"
	"github.com/agustin-sarasua/pimbay/app/user"
)

var (
	StorageBucket     *storage.BucketHandle
	StorageBucketName string
)

func init() {
	fmt.Println("Running init...")
	os.Setenv("DATASTORE_DATASET", "pimbay-accounting")
	os.Setenv("DATASTORE_EMULATOR_HOST", "localhost:8081")
	os.Setenv("DATASTORE_EMULATOR_HOST_PATH", "localhost:8081/datastore")
	os.Setenv("DATASTORE_HOST", "http://localhost:8081")
	os.Setenv("DATASTORE_PROJECT_ID", "pimbay-accounting")
	os.Setenv("GCLOUD_STORAGE_BUCKET", "pimbay-accounting.appspot.com")

	account.AccountDB, _ = configureAccountDatastoreDB("pimbay-accounting")
	user.UserDB, _ = configureUserDatastoreDB("pimbay-accounting")
	// [START storage]
	// To configure Cloud Storage, uncomment the following lines and update the
	// bucket name.
	//
	StorageBucketName = "pimbay-accounting.appspot.com"
	StorageBucket, _ = configureStorage(StorageBucketName)
	// [END storage]
	api.FbAPI = api.NewFirebaseAPI()
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()
}

func configureAccountDatastoreDB(projectID string) (account.AccountDatabase, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return account.NewDatastoreDB(ctx, client)
}

func configureUserDatastoreDB(projectID string) (user.UserDatabase, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return user.NewDatastoreDB(ctx, client)
}

func configureStorage(bucketID string) (*storage.BucketHandle, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.Bucket(bucketID), nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}
