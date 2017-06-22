package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/agustin-sarasua/pimbay/api"
)

const (
	signInEndpoint         = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=AIzaSyAkR4u8iQBLckmYNtWSx9fmJNSilyWc__A"
	signUpEndpoint         = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/signupNewUser?key=AIzaSyAkR4u8iQBLckmYNtWSx9fmJNSilyWc__A"
	getAccountInfoEndpoint = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/getAccountInfo?key=AIzaSyAkR4u8iQBLckmYNtWSx9fmJNSilyWc__A"
	applicationContent     = "application/json"
)

func SignupNewUser(email, pwd string) api.SignUpResponse {
	fmt.Println("SignUp new User")
	jsonValue, _ := json.Marshal(api.SignUpRequest{Email: email, Password: pwd, ReturnSecureToken: true})
	resp, err := http.Post(signUpEndpoint, applicationContent, bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}

	fmt.Println("Unmarshalling response")
	var r api.SignUpResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		panic(err)
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
