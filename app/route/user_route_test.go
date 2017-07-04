package route_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"golang.org/x/net/context"

	"github.com/agustin-sarasua/pimbay"
	"github.com/agustin-sarasua/pimbay/app/api"
	"github.com/agustin-sarasua/pimbay/app/model"
	"github.com/agustin-sarasua/pimbay/app/route"
)

func TestSignupNewUser(t *testing.T) {
	testEmail := "t@t.com"
	jsonValue, _ := json.Marshal(api.SignupUserRestMsg{Name: "Agustin", LastName: "Sarasua", Email: testEmail, Password: "pwdTest1234"})
	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Failed to create req1: %v", err)
	}
	c1 := context.Background()

	handler := http.HandlerFunc(route.SignupNewUserEndpoint)
	req = req.WithContext(c1)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	r, _ := pimbay.DB.GetUserByEmail(c1, testEmail)
	if r == nil {
		t.Errorf("The user has not signed up")
	}
	pimbay.DB.Cleanup()
	r, err = pimbay.DB.GetUserByEmail(c1, testEmail)
	if err != nil {
		t.Fatalf("Failed to getting user: %v", err)
	}
	fmt.Println(r)
}

func TestGetUser(t *testing.T) {
	id, err := pimbay.DB.SaveUser(context.Background(), &model.User{ID: 1234, FirebaseID: "asdf", Email: "agustinsarasua@gmail.com", Name: "Agustin"})
	rr := httptest.NewRecorder()
	var buffer bytes.Buffer
	buffer.WriteString("/user/")
	buffer.WriteString(strconv.FormatInt(id, 10))
	fmt.Println(buffer.String())
	req, err := http.NewRequest("GET", buffer.String(), nil)
	if err != nil {
		t.Fatalf("Failed to create req1: %v", err)
	}
	r := startTestServer()
	r.ServeHTTP(rr, req)
	var res *model.User
	json.NewDecoder(rr.Body).Decode(&res)
	fmt.Print(rr.Body.String())
	pimbay.DB.Cleanup()
	if res == nil {
		t.Fatalf("Failed to get User: %v", err)
	}

}
