package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/appengine"

	"cloud.google.com/go/datastore"

	"github.com/agustin-sarasua/pimbay/model"
	"github.com/agustin-sarasua/pimbay/web"
	"github.com/golang/glog"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	DB model.UserDatabase
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()
}

func main() {
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.Print("Hello, log file!")

	//DB, _ = model.NewMongoDB("localhost", "gopimbay", "gopimbay", "gopimbay", "user")
	DB, _ = configureDatastoreDB("pimbay-accounting")
	defer DB.Close()

	//Fin Mongo
	defer glog.Flush()

	//Gorilla MUX
	router := mux.NewRouter()

	router.HandleFunc("/signin", use(web.SigninUserEndpoint, web.BasicAuth)).Methods("POST")
	router.HandleFunc("/signup", web.SignupNewUserEndpoint(DB)).Methods("POST")

	router.HandleFunc("/account", use(web.CreateAccountEndpoint(DB), web.ValidateToken)).Methods("POST")

	router.HandleFunc("/hello", use(web.GetAccountInfo, web.ValidateToken)).Methods("GET")

	router.Methods("GET").Path("/_ah/health").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))
	appengine.Main()
	fmt.Println("Hello there...")
	//log.Fatal(http.ListenAndServe(":12345", router))
}

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

func configureDatastoreDB(projectID string) (model.UserDatabase, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return model.NewDatastoreDB(client)
}
