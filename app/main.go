package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/appengine"

	"cloud.google.com/go/storage"

	"github.com/agustin-sarasua/pimbay/app/account"
	"github.com/agustin-sarasua/pimbay/app/reports"
	"github.com/agustin-sarasua/pimbay/app/route"
	"github.com/agustin-sarasua/pimbay/app/user"
	"github.com/golang/glog"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	storageClient *storage.Client

	// Set this in app.yaml when running in production.
	bucket = os.Getenv("GCLOUD_STORAGE_BUCKET")
)

func main() {
	fmt.Println("Running main...")
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.Print("Hello, log file!")

	var err error
	storageClient, err = storage.NewClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer user.UserDB.Close()

	defer glog.Flush()
	StartServer()
	fmt.Println("Hello there...")
	//log.Fatal(http.ListenAndServe(":12345", router))
}

func StartServer() {
	//Gorilla MUX
	router := mux.NewRouter()

	router.HandleFunc("/signin", use(user.SigninUserEndpoint, route.BasicAuth)).Methods("POST")
	router.HandleFunc("/signup", user.SignupNewUserEndpoint).Methods("POST")
	router.HandleFunc("/user/{id:[0-9]+}", user.GetUser).Methods("GET")

	router.HandleFunc("/user/{id:[0-9]+}/accounts", user.GetUser).Methods("GET")

	router.HandleFunc("/account", use(account.CreateAccountEndpoint, route.ValidateToken)).Methods("POST")

	router.HandleFunc("/hello", use(user.GetAccountInfo, route.ValidateToken)).Methods("GET")

	router.Methods("GET").Path("/_ah/health").HandlerFunc(route.HealthCheckHandler)

	router.HandleFunc("/account/{id:[0-9]+}/report", use(reports.SaveAccountReportEndpoint, route.ValidateToken)).Methods("POST")
	router.HandleFunc("/report/{name:.*}/process", use(reports.SaveAccountReportEndpoint, route.ValidateToken)).Methods("POST")

	router.HandleFunc("/report/upload", reports.FormHandler)
	router.HandleFunc("/upload", reports.UploadHandler)

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))
	appengine.Main()
}

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}
