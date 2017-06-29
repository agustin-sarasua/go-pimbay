package pimbay

import (
	"context"
	"flag"
	"fmt"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/agustin-sarasua/pimbay/api"
	"github.com/agustin-sarasua/pimbay/db"
)

var (
	DB    db.Database
	FbAPI api.FirebaseAPI
)

func init() {
	fmt.Println("Running init...")
	DB, _ = configureDatastoreDB("pimbay-accounting")
	FbAPI = api.NewFirebaseAPI()
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()
}

func configureDatastoreDB(projectID string) (db.Database, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return db.NewDatastoreDB(ctx, client)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}
