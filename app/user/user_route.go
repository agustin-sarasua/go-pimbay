package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang.org/x/net/context"

	"github.com/agustin-sarasua/pimbay/app/util"
	"github.com/gorilla/mux"

	"fmt"
)

func GetUser(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(req)["id"], 10, 64)
	fmt.Printf("Getting User %d...", id)
	if err != nil {
		fmt.Errorf("bad user id: %v", err)
		return
	}
	u, err := UserDB.GetUser(context.Background(), id)
	fmt.Println(u)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func SignupNewUserEndpoint(w http.ResponseWriter, req *http.Request) {
	var msg SignupUserRestMsg
	err := json.NewDecoder(req.Body).Decode(&msg)

	if err != nil {
		util.ErrorWithJSON(w, "", http.StatusBadRequest)
		return
	}

	SignupNewUser(UserDB, &msg)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)

}

func SigninUserEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("User Logged in with token: ", w.Header().Get("token"))

}

func GetAccountInfo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	uid := ctx.Value(999)
	fmt.Println("User Id: ", uid)
	fmt.Println("Hello")
}
