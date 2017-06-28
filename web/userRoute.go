package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/agustin-sarasua/pimbay/api"
	"github.com/agustin-sarasua/pimbay/db"
	"github.com/gorilla/mux"

	"fmt"

	"github.com/agustin-sarasua/pimbay/service"
)

func GetUser(db db.UserDatabase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		id, err := strconv.ParseInt(mux.Vars(req)["id"], 10, 64)
		if err != nil {
			fmt.Errorf("bad user id: %v", err)
			return
		}
		u, err := db.GetUser(id)

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(u)
	}
}

func SignupNewUserEndpoint(db db.UserDatabase) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var msg api.SignupUserRestMsg
		err := json.NewDecoder(req.Body).Decode(&msg)

		if err != nil {
			ErrorWithJSON(w, "", http.StatusBadRequest)
			return
		}

		service.SignupNewUser(db, &msg)
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(msg)
	}
}

func SigninUserEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("User Logged in with token: ", w.Header().Get("token"))

}

func GetAccountInfo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello")
}
