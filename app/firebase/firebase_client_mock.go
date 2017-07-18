package firebase

import "fmt"

type firebaseMockedAPI struct{}

func (f *firebaseMockedAPI) Signin() {
	fmt.Println("Mocked FirebaseAPI")
}

func (f *firebaseMockedAPI) Signup(e, p string, rs bool) (*SignUpResponse, error) {
	return &SignUpResponse{}, nil
}

func NewFirebaseMockedAPI() FirebaseAPI {
	return &firebaseMockedAPI{}
}
