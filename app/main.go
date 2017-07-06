package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"google.golang.org/appengine"

	"cloud.google.com/go/storage"

	"github.com/GoogleCloudPlatform/golang-samples/getting-started/bookshelf"
	"github.com/agustin-sarasua/pimbay"
	"github.com/agustin-sarasua/pimbay/app/route"
	"github.com/agustin-sarasua/pimbay/app/user"
	"github.com/golang/glog"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
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

	router.HandleFunc("/account", use(route.CreateAccountEndpoint, route.ValidateToken)).Methods("POST")

	router.HandleFunc("/hello", use(user.GetAccountInfo, route.ValidateToken)).Methods("GET")

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

func uploadHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(rw, "", http.StatusMethodNotAllowed)
		return
	}

	//ctx := appengine.NewContext(r)
	// ctx := context.Background()

	f, fh, err := r.FormFile("file")
	if err == http.ErrMissingFile {
		return
	}
	if err != nil {
		return
	}

	if pimbay.StorageBucket == nil {
		return
	}

	// random filename, retaining existing extension.
	name := uuid.NewV4().String() + path.Ext(fh.Filename)

	ctx := context.Background()
	w := pimbay.StorageBucket.Object(name).NewWriter(ctx)
	w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	w.ContentType = fh.Header.Get("Content-Type")

	// Entries are immutable, be aggressive about caching (1 day).
	w.CacheControl = "public, max-age=86400"

	if _, err := io.Copy(w, f); err != nil {
		return
	}
	if err := w.Close(); err != nil {
		return
	}

	const publicURL = "https://storage.googleapis.com/%s/%s"

	// f, fh, err := r.FormFile("file")
	// if err != nil {
	// 	msg := fmt.Sprintf("Could not get file: %v", err)
	// 	http.Error(w, msg, http.StatusBadRequest)
	// 	return
	// }
	// defer f.Close()

	// fmt.Println(bucket)

	// sw := storageClient.Bucket(bucket).Object(fh.Filename).NewWriter(ctx)
	// if _, err := io.Copy(sw, f); err != nil {
	// 	msg := fmt.Sprintf("Could not write file: %v", err)
	// 	http.Error(w, msg, http.StatusInternalServerError)
	// 	return
	// }

	// if err := sw.Close(); err != nil {
	// 	msg := fmt.Sprintf("Could not put file: %v", err)
	// 	http.Error(w, msg, http.StatusInternalServerError)
	// 	return
	// }

	// u, _ := url.Parse("/" + bucket + "/" + sw.Attrs().Name)

	fmt.Fprintf(rw, fmt.Sprintf(publicURL, pimbay.StorageBucketName, name))
}

// uploadFileFromForm uploads a file if it's present in the "image" form field.
func uploadFileFromForm(r *http.Request) (url string, err error) {
	f, fh, err := r.FormFile("image")
	if err == http.ErrMissingFile {
		return "", nil
	}
	if err != nil {
		return "", err
	}

	if pimbay.StorageBucket == nil {
		return "", errors.New("storage bucket is missing - check config.go")
	}

	// random filename, retaining existing extension.
	name := uuid.NewV4().String() + path.Ext(fh.Filename)

	ctx := context.Background()
	w := bookshelf.StorageBucket.Object(name).NewWriter(ctx)
	w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	w.ContentType = fh.Header.Get("Content-Type")

	// Entries are immutable, be aggressive about caching (1 day).
	w.CacheControl = "public, max-age=86400"

	if _, err := io.Copy(w, f); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}

	const publicURL = "https://storage.googleapis.com/%s/%s"
	return fmt.Sprintf(publicURL, bookshelf.StorageBucketName, name), nil
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
