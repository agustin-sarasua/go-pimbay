package email

import (
	"bytes"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func IncomingMail(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	defer r.Body.Close()
	var b bytes.Buffer
	if _, err := b.ReadFrom(r.Body); err != nil {
		log.Errorf(ctx, "Error reading body: %v", err)
		return
	}
	log.Infof(ctx, "Received mail: %v", b)
}
