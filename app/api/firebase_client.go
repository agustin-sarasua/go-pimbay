package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type firebaseAPI struct{}

func (f *firebaseAPI) Signin() {
	fmt.Println("Real FirebaseAPI")
}

func (f *firebaseAPI) Signup(e, p string, rs bool) *SignUpResponse {
	jsonValue, _ := json.Marshal(SignUpRequest{Email: e, Password: p, ReturnSecureToken: rs})
	resp, err := http.Post(signUpEndpoint, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	var r SignUpResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		panic(err)
	}
	return &r
}

func NewFirebaseAPI() FirebaseAPI {
	return &firebaseAPI{}
}
