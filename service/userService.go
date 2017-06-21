package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/agustin-sarasua/pimbay/api"
)

const (
	signUpEndpoint     = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/signupNewUser?key=AIzaSyAkR4u8iQBLckmYNtWSx9fmJNSilyWc__A"
	applicationContent = "application/json"
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
