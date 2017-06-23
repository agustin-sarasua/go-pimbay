package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"time"

	"github.com/agustin-sarasua/pimbay/api"
	"github.com/agustin-sarasua/pimbay/model"
)

const (
	signInEndpoint         = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=AIzaSyAkR4u8iQBLckmYNtWSx9fmJNSilyWc__A"
	signUpEndpoint         = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/signupNewUser?key=AIzaSyAkR4u8iQBLckmYNtWSx9fmJNSilyWc__A"
	getAccountInfoEndpoint = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/getAccountInfo?key=AIzaSyAkR4u8iQBLckmYNtWSx9fmJNSilyWc__A"
	applicationContent     = "application/json"
	dbName                 = "gopimbay"
	userCollection         = "user"
)

func SignupNewUser(s *mgo.Session, m *api.SignupUserRestMsg) api.SignUpResponse {
	fmt.Println("SignUp new User")
	jsonValue, _ := json.Marshal(api.SignUpRequest{Email: m.Email, Password: m.Password, ReturnSecureToken: true})
	resp, err := http.Post(signUpEndpoint, applicationContent, bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}

	var r api.SignUpResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		panic(err)
	}
	u := model.User{ID: r.LocalID, Name: m.Name, LastName: m.LastName, CreatedDate: time.Now(), Birthdate: m.Birthdate, Email: m.Email, Sex: m.Sex}
	ch := saveUser(s, &u)
	ok := <-ch
	if !ok {
		fmt.Println("Could not save user")
	}
	fmt.Println(r)
	return r
}

func SigninUser(email, pwd string) <-chan *api.SignUpResponse {
	out := make(chan *api.SignUpResponse)
	go func() {
		fmt.Println("SignIn User ", email)
		jsonValue, _ := json.Marshal(api.SignUpRequest{Email: email, Password: pwd, ReturnSecureToken: true})
		resp, err := http.Post(signInEndpoint, applicationContent, bytes.NewBuffer(jsonValue))
		if err != nil {
			panic(err)
		}

		var r api.SignUpResponse
		err = json.NewDecoder(resp.Body).Decode(&r)
		if err != nil {
			panic(err)
		}
		out <- &r
	}()
	return out
}

func GetAccountInfo(tkn string) <-chan *api.AccountInfoReponse {
	out := make(chan *api.AccountInfoReponse)
	go func() {
		values := map[string]string{"idToken": tkn}
		jsonValue, _ := json.Marshal(values)
		resp, err := http.Post(getAccountInfoEndpoint, applicationContent, bytes.NewBuffer(jsonValue))
		if err != nil {
			panic(err)
		}

		var r api.AccountInfoReponse
		err = json.NewDecoder(resp.Body).Decode(&r)
		if err != nil {
			panic(err)
		}
		out <- &r
	}()
	return out
}

func saveUser(s *mgo.Session, u *model.User) <-chan bool {
	out := make(chan bool)
	//defer finally()
	go func() {
		session := s.Copy()
		defer session.Close()

		c := session.DB(dbName).C(userCollection)
		err := c.Insert(u)
		if err != nil {
			if mgo.IsDup(err) {
				panic(err)
			}
			log.Println("Failed insert User: ", err)
			out <- false
		} else {
			out <- true
		}
	}()
	return out
}
