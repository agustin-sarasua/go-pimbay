package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/agustin-sarasua/pimbay/web"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
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
	// Mongo
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost"},
		Timeout:  60 * time.Second,
		Database: "gopimbay",
		Username: "gopimbay",
		Password: "gopimbay",
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	//ensureIndex(session)
	//Fin Mongo
	defer glog.Flush()

	//Gorilla MUX
	router := mux.NewRouter()

	router.HandleFunc("/signin", use(web.SignupNewUserEndpoint(session), web.BasicAuth)).Methods("GET")
	router.HandleFunc("/hello", use(web.GetAccountInfo, web.ValidateToken)).Methods("GET")
	//router.HandleFunc("/fraud-info", web.ProcessSaveFraudInfo(session)).Methods("POST")

	//var p = model.Person{Id: 1, Name: "Agustin", LastName: "Sarasua", Address: nil}
	//people = append(people, p)
	fmt.Println("Hello there")
	log.Fatal(http.ListenAndServe(":12345", router))
}

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}
