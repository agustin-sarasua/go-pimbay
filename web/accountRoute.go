package web

import (
	"encoding/json"
	"net/http"

	"time"

	"github.com/agustin-sarasua/pimbay/model"
	"github.com/agustin-sarasua/pimbay/service"
	mgo "gopkg.in/mgo.v2"
)

func CreateAccountEndpoint(s *mgo.Session) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		session := s.Copy()
		defer session.Close()

		var msg model.Account
		err := json.NewDecoder(req.Body).Decode(&msg)

		if err != nil {
			ErrorWithJSON(w, "", http.StatusBadRequest)
			return
		}
		msg.CreatedDate = time.Now()
		service.SaveAccount(s, &msg)
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(msg)
	}
}
