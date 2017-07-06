package route_test

import (
	"fmt"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/agustin-sarasua/pimbay"
	"github.com/agustin-sarasua/pimbay/app/api"
	"github.com/agustin-sarasua/pimbay/app/db"
	"github.com/agustin-sarasua/pimbay/app/route"
	"github.com/agustin-sarasua/pimbay/app/user"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

func init() {
	fmt.Println("Running init test...")
	os.Setenv("DATASTORE_DATASET", "pimbay-accounting")
	os.Setenv("DATASTORE_EMULATOR_HOST", "localhost:8081")
	os.Setenv("DATASTORE_EMULATOR_HOST_PATH", "localhost:8081/datastore")
	os.Setenv("DATASTORE_HOST", "http://localhost:8081")
	os.Setenv("DATASTORE_PROJECT_ID", "pimbay-accounting")
	pimbay.FbAPI = api.NewFirebaseMockedAPI()
	pimbay.DB, _ = configureDatastoreDB("pimbay-accounting")
}

func startTestServer() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/user/{id:[0-9]+}", user.GetUser).Methods("GET")
	router.HandleFunc("/user/{id:[0-9]+}/accounts", user.GetUser).Methods("GET")
	router.Methods("GET").Path("/_ah/health").HandlerFunc(route.HealthCheckHandler)
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))
	return router
}

func configureDatastoreDB(projectID string) (db.Database, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return db.NewDatastoreDB(ctx, client)
}
