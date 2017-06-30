package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/agustin-sarasua/pimbay"
	"github.com/agustin-sarasua/pimbay/app/route"
	"github.com/golang/glog"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	fmt.Println("Running main...")
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.Print("Hello, log file!")

	//DB, _ = model.NewMongoDB("localhost", "gopimbay", "gopimbay", "gopimbay", "user")

	defer pimbay.DB.Close()

	defer glog.Flush()

	//Gorilla MUX
	router := mux.NewRouter()

	router.HandleFunc("/signin", use(route.SigninUserEndpoint, route.BasicAuth)).Methods("POST")
	router.HandleFunc("/signup", route.SignupNewUserEndpoint).Methods("POST")
	router.HandleFunc("/user/{id:[0-9]+}", route.GetUser).Methods("GET")

	router.HandleFunc("/user/{id:[0-9]+}/accounts", route.GetUser).Methods("GET")

	router.HandleFunc("/account", use(route.CreateAccountEndpoint, route.ValidateToken)).Methods("POST")

	router.HandleFunc("/hello", use(route.GetAccountInfo, route.ValidateToken)).Methods("GET")

	router.Methods("GET").Path("/_ah/health").HandlerFunc(route.HealthCheckHandler)

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
