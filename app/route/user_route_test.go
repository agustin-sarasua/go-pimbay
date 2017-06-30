package route_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/net/context"

	"github.com/agustin-sarasua/pimbay"
	"github.com/agustin-sarasua/pimbay/app/api"
	"github.com/agustin-sarasua/pimbay/app/main"
)

func TestSignupNewUser(t *testing.T) {
	testEmail := "t@t.com"
	jsonValue, _ := json.Marshal(api.SignupUserRestMsg{Name: "Agustin", LastName: "Sarasua", Email: testEmail, Password: "pwdTest1234"})
	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Failed to create req1: %v", err)
	}
	c1 := context.Background()

	handler := http.HandlerFunc(main.SignupNewUserEndpoint)
	req = req.WithContext(c1)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	r, _ := pimbay.DB.GetUserByEmail(c1, testEmail)
	if r == nil {
		t.Errorf("The user has not signed up")
	}
	err = pimbay.DB.DeleteUser(r.ID)
	if err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}
	r, err = pimbay.DB.GetUserByEmail(c1, testEmail)
	if err != nil {
		t.Fatalf("Failed to getting user: %v", err)
	}
	fmt.Println(r)
}

func init() {
	fmt.Println("Running init test...")
	pimbay.FbAPI = api.NewFirebaseMockedAPI()

	//pimbay.DB =
}
