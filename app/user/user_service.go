package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/agustin-sarasua/pimbay/app/firebase"
)

const (
	signInEndpoint         = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=AIzaSyAkR4u8iQBLckmYNtWSx9fmJNSilyWc__A"
	getAccountInfoEndpoint = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/getAccountInfo?key=AIzaSyAkR4u8iQBLckmYNtWSx9fmJNSilyWc__A"
	userCollection         = "user"
	applicationContent     = "application/json"
)

func SignupNewUser(db UserDatabase, m *SignupUserRestMsg) *firebase.SignUpResponse {
	fmt.Println("SignUp new User")
	r, _ := firebase.FbAPI.Signup(m.Email, m.Password, true)

	u := User{FirebaseID: r.LocalID, Name: m.Name, LastName: m.LastName, CreatedDate: time.Now(), Birthdate: m.Birthdate, Email: m.Email, Sex: m.Sex}
	id, _ := db.SaveUser(context.Background(), &u)
	u.ID = id
	return r
}

func SigninUser(email, pwd string) <-chan *firebase.SignUpResponse {
	out := make(chan *firebase.SignUpResponse)
	go func() {
		fmt.Println("SignIn User ", email, pwd)
		jsonValue, _ := json.Marshal(firebase.SignUpRequest{Email: email, Password: pwd, ReturnSecureToken: true})
		resp, err := http.Post(signInEndpoint, applicationContent, bytes.NewBuffer(jsonValue))
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)

		var r firebase.SignUpResponse
		err = json.NewDecoder(resp.Body).Decode(&r)
		if err != nil {
			panic(err)
		}
		out <- &r
	}()
	return out
}

func GetAccountInfoS(tkn string) <-chan *firebase.AccountInfoReponse {
	out := make(chan *firebase.AccountInfoReponse)
	go func() {
		values := map[string]string{"idToken": tkn}
		jsonValue, _ := json.Marshal(values)
		resp, err := http.Post(getAccountInfoEndpoint, applicationContent, bytes.NewBuffer(jsonValue))
		if err != nil {
			panic(err)
		}

		var r firebase.AccountInfoReponse
		err = json.NewDecoder(resp.Body).Decode(&r)
		if err != nil {
			panic(err)
		}
		out <- &r
	}()
	return out
}
