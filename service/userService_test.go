package service_test

import (
	"fmt"
	"testing"

	"github.com/agustin-sarasua/pimbay/service"
)

func TestSignupNewUser(t *testing.T) {
	fmt.Println("Running test")
	service.SignupNewUser("agustinsarasua@gmail.com", "testPwd")
}
