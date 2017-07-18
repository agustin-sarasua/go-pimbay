package firebase

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

func (f *firebaseAPI) Signup(e, p string, rs bool) (*SignUpResponse, error) {
	jsonValue, _ := json.Marshal(SignUpRequest{Email: e, Password: p, ReturnSecureToken: rs})
	// If you don't read the response body you still need to close it
	resp, err := http.Post(signUpEndpoint, "application/json", bytes.NewBuffer(jsonValue))
	if resp != nil {
		//Always Close the responseBody
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var r SignUpResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		panic(err)
	}
	return &r, err
}

func NewFirebaseAPI() FirebaseAPI {
	return &firebaseAPI{}
}
