package pimbay

import (
	"context"
	"flag"
	"fmt"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/agustin-sarasua/pimbay/app/api"
	"github.com/agustin-sarasua/pimbay/app/db"
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
	// export DATASTORE_DATASET=pimbay-accounting
	// export DATASTORE_EMULATOR_HOST=localhost:8081
	// export DATASTORE_EMULATOR_HOST_PATH=localhost:8081/datastore
	// export DATASTORE_HOST=http://localhost:8081
	// export DATASTORE_PROJECT_ID=pimbay-accounting
	os.Setenv("DATASTORE_DATASET", "pimbay-accounting")
	os.Setenv("DATASTORE_EMULATOR_HOST", "localhost:8081")
	os.Setenv("DATASTORE_EMULATOR_HOST_PATH", "localhost:8081/datastore")
	os.Setenv("DATASTORE_HOST", "http://localhost:8081")
	os.Setenv("DATASTORE_PROJECT_ID", "pimbay-accounting")

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
