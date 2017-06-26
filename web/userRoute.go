package web

import (
	"encoding/json"
	"net/http"

	"github.com/agustin-sarasua/pimbay/api"
	"github.com/agustin-sarasua/pimbay/model"

	"fmt"

	"github.com/agustin-sarasua/pimbay/service"
)

func SignupNewUserEndpoint(db model.UserDatabase) func(w http.ResponseWriter, req *http.Request) {
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
