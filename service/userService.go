package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/agustin-sarasua/pimbay"
	"github.com/agustin-sarasua/pimbay/api"
	"github.com/agustin-sarasua/pimbay/db"
	"github.com/agustin-sarasua/pimbay/model"
)

const (
	signInEndpoint         = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=AIzaSyAkR4u8iQBLckmYNtWSx9fmJNSilyWc__A"
	signUpEndpoint         = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/signupNewUser?key=AIzaSyAkR4u8iQBLckmYNtWSx9fmJNSilyWc__A"
	getAccountInfoEndpoint = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/getAccountInfo?key=AIzaSyAkR4u8iQBLckmYNtWSx9fmJNSilyWc__A"
	userCollection         = "user"
)

func SignupNewUser(db db.Database, m *api.SignupUserRestMsg) api.SignUpResponse {
	fmt.Println("SignUp new User")
	pimbay.FbAPI.Signin()
	jsonValue, _ := json.Marshal(api.SignUpRequest{Email: m.Email, Password: m.Password, ReturnSecureToken: true})
	resp, err := http.Post(signUpEndpoint, applicationContent, bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	var r api.SignUpResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		panic(err)
	}
	u := model.User{FirebaseID: r.LocalID, Name: m.Name, LastName: m.LastName, CreatedDate: time.Now(), Birthdate: m.Birthdate, Email: m.Email, Sex: m.Sex}
	id, _ := db.SaveUser(context.Background(), &u)
	u.ID = id
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
