package api

import "fmt"

type firebaseMockedAPI struct{}

func (f *firebaseMockedAPI) Signin() {
	fmt.Println("Mocked FirebaseAPI")
}

func (f *firebaseMockedAPI) Signup(u, p string, rs bool) *SignUpResponse {
	return &SignUpResponse{}
}

func NewFirebaseMockedAPI() FirebaseAPI {
	return &firebaseMockedAPI{}
}
