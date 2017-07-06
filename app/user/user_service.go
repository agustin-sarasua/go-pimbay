package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/agustin-sarasua/pimbay/app/api"
)

const (
	signInEndpoint         = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=AIzaSyAkR4u8iQBLckmYNtWSx9fmJNSilyWc__A"
	getAccountInfoEndpoint = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/getAccountInfo?key=AIzaSyAkR4u8iQBLckmYNtWSx9fmJNSilyWc__A"
	userCollection         = "user"
	applicationContent     = "application/json"
)

func SignupNewUser(db UserDatabase, m *api.SignupUserRestMsg) *api.SignUpResponse {
	fmt.Println("SignUp new User")
	r, _ := api.FbAPI.Signup(m.Email, m.Password, true)

	u := User{FirebaseID: r.LocalID, Name: m.Name, LastName: m.LastName, CreatedDate: time.Now(), Birthdate: m.Birthdate, Email: m.Email, Sex: m.Sex}
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

func GetAccountInfoS(tkn string) <-chan *api.AccountInfoReponse {
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
