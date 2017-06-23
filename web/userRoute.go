package web

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"

	"github.com/agustin-sarasua/pimbay/model"
	mgo "gopkg.in/mgo.v2"
)

func SignupNewUserEndpoint(s *mgo.Session) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		session := s.Copy()
		defer session.Close()

		//params := mux.Vars(req)
		var user model.User
		err := json.NewDecoder(req.Body).Decode(&user)
		req.Header.Get("Authorization")
		//user.Id, _ = strconv.Atoi(params["id"])
		//service.SignupNewUser()
		c := session.DB("gopimbay").C("user")
		err = c.Insert(user)
		if err != nil {
			if mgo.IsDup(err) {
				ErrorWithJSON(w, "Book with this ISBN already exists", http.StatusBadRequest)
				return
			}

			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed insert book: ", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		//w.Header().Set("Location", r.URL.Path+"/"+book.ISBN)
		w.WriteHeader(http.StatusCreated)

		// people = append(people, person)*/
		json.NewEncoder(w).Encode(user)
	}
}

func SigninUserEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("User Logged in with token: ", w.Header().Get("token"))

}

func GetAccountInfo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello")
}
