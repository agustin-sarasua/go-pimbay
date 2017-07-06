package account

import (
	"encoding/json"
	"net/http"

	"time"

	"github.com/agustin-sarasua/pimbay/app/util"
)

func CreateAccountEndpoint(w http.ResponseWriter, req *http.Request) {
	var msg Account
	err := json.NewDecoder(req.Body).Decode(&msg)

	if err != nil {
		util.ErrorWithJSON(w, "", http.StatusBadRequest)
		return
	}
	msg.CreatedDate = time.Now()
	//service.CreateAccount(s, &msg)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}

func ListUserAccountsEndpoint(rw http.ResponseWriter, r *http.Request) {

}
