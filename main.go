package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/agustin-sarasua/pimbay/model"
	"github.com/agustin-sarasua/pimbay/web"
	"github.com/golang/glog"
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

	DB, _ = model.NewMongoDB("localhost", "gopimbay", "gopimbay", "gopimbay", "user")
	//DB, err = configureDatastoreDB()

	//Fin Mongo
	defer glog.Flush()

	//Gorilla MUX
	router := mux.NewRouter()

	router.HandleFunc("/signin", use(web.SigninUserEndpoint, web.BasicAuth)).Methods("POST")
	router.HandleFunc("/signup", web.SignupNewUserEndpoint(DB)).Methods("POST")

	router.HandleFunc("/account", use(web.CreateAccountEndpoint(DB), web.ValidateToken)).Methods("POST")

	router.HandleFunc("/hello", use(web.GetAccountInfo, web.ValidateToken)).Methods("GET")

	fmt.Println("Hello there...")
	log.Fatal(http.ListenAndServe(":12345", router))
}

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}
