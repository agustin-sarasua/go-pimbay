package route

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang.org/x/net/context"

	"github.com/agustin-sarasua/pimbay"
	"github.com/agustin-sarasua/pimbay/app/api"
	"github.com/gorilla/mux"

	"fmt"

	"github.com/agustin-sarasua/pimbay/app/service"
)

func GetUser(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(req)["id"], 10, 64)
	if err != nil {
		fmt.Errorf("bad user id: %v", err)
		return
	}
	u, err := pimbay.DB.GetUser(context.Background(), id)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func SignupNewUserEndpoint(w http.ResponseWriter, req *http.Request) {
	var msg api.SignupUserRestMsg
	err := json.NewDecoder(req.Body).Decode(&msg)

	if err != nil {
		ErrorWithJSON(w, "", http.StatusBadRequest)
		return
	}

	service.SignupNewUser(pimbay.DB, &msg)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)

}

func SigninUserEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("User Logged in with token: ", w.Header().Get("token"))

}

func GetAccountInfo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello")
}
