package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"google.golang.org/appengine"

	"cloud.google.com/go/storage"

	"github.com/agustin-sarasua/pimbay"
	"github.com/agustin-sarasua/pimbay/app/route"
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

	defer pimbay.DB.Close()

	defer glog.Flush()
	StartServer()
	fmt.Println("Hello there...")
	//log.Fatal(http.ListenAndServe(":12345", router))
}

func StartServer() {
	//Gorilla MUX
	router := mux.NewRouter()

	router.HandleFunc("/signin", use(route.SigninUserEndpoint, route.BasicAuth)).Methods("POST")
	router.HandleFunc("/signup", route.SignupNewUserEndpoint).Methods("POST")
	router.HandleFunc("/user/{id:[0-9]+}", route.GetUser).Methods("GET")

	router.HandleFunc("/user/{id:[0-9]+}/accounts", route.GetUser).Methods("GET")

	router.HandleFunc("/account", use(route.CreateAccountEndpoint, route.ValidateToken)).Methods("POST")

	router.HandleFunc("/hello", use(route.GetAccountInfo, route.ValidateToken)).Methods("GET")

	router.Methods("GET").Path("/_ah/health").HandlerFunc(route.HealthCheckHandler)

	router.HandleFunc("/report/upload", formHandler)
	router.HandleFunc("/upload", uploadHandler)

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))
	appengine.Main()
}

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	//ctx := appengine.NewContext(r)
	ctx := context.Background()
	f, fh, err := r.FormFile("file")
	if err != nil {
		msg := fmt.Sprintf("Could not get file: %v", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	defer f.Close()

	fmt.Println(bucket)

	sw := storageClient.Bucket(bucket).Object(fh.Filename).NewWriter(ctx)
	if _, err := io.Copy(sw, f); err != nil {
		msg := fmt.Sprintf("Could not write file: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	if err := sw.Close(); err != nil {
		msg := fmt.Sprintf("Could not put file: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	u, _ := url.Parse("/" + bucket + "/" + sw.Attrs().Name)

	fmt.Fprintf(w, "Successful! URL: https://storage.googleapis.com%s", u.EscapedPath())
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, formHTML)
}

const formHTML = `<!DOCTYPE html>
<html>
  <head>
    <title>Storage</title>
    <meta charset="utf-8">
  </head>
  <body>
    <form method="POST" action="/upload" enctype="multipart/form-data">
      <input type="file" name="file">
      <input type="submit">
    </form>
  </body>
</html>`
