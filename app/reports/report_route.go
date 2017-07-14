package reports

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/agustin-sarasua/pimbay/app/route"
	"github.com/agustin-sarasua/pimbay/app/util"
	"github.com/gorilla/mux"
)

var (
	StorageBucket     *storage.BucketHandle
	StorageBucketName string
)

func SaveAccountReportEndpoint(w http.ResponseWriter, req *http.Request) {
	uid := route.GetUserID(req)
	if uid == 0 {
		util.ErrorWithJSON(w, "UserID is 0", http.StatusBadRequest)
		return
	}

	var msg SaveReportRequest
	err := json.NewDecoder(req.Body).Decode(&msg)

	if err != nil {
		util.ErrorWithJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	// SignupNewUser(UserDB, &msg)
	// w.Header().Set("Content-Type", "application/json")

	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(msg)
}

func ProcessAccountReportEndpoint(w http.ResponseWriter, req *http.Request) {
	uid := route.GetUserID(req)

	if uid == 0 {
		util.ErrorWithJSON(w, "UserID is 0", http.StatusBadRequest)
		return
	}

	_, err := mux.Vars(req)["name"]
	if err {
		util.ErrorWithJSON(w, "", http.StatusBadRequest)
		return
	}

	// SignupNewUser(UserDB, &msg)
	// w.Header().Set("Content-Type", "application/json")

	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(msg)
}

func UploadHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(rw, "", http.StatusMethodNotAllowed)
		return
	}

	f, fh, err := r.FormFile("file")

	if err == http.ErrMissingFile {
		return
	}
	if err != nil {
		return
	}

	nc := PushReportToCloudStorage(f, fh)
	rc := ProcessReport(f, fh)
	n := <-nc
	rs := <-rc

	fmt.Println(rs)

	const publicURL = "https://storage.googleapis.com/%s/%s"

	fmt.Fprintf(rw, fmt.Sprintf(publicURL, StorageBucketName, n))
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
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
