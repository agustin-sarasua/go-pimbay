package route_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/agustin-sarasua/pimbay"
	"github.com/agustin-sarasua/pimbay/app/api"
	"github.com/agustin-sarasua/pimbay/app/model"
	"github.com/agustin-sarasua/pimbay/app/route"
)

func TestCreateAccountEndpoint(t *testing.T) {
	testEmail := "t@t.com"
	jsonValue, _ := json.Marshal(model.Account{Name: "Itau"})
	req, err := http.NewRequest("POST", "/account", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Failed to create req1: %v", err)
	}
	handler := http.HandlerFunc(route.CreateAccountEndpoint)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	err = json.NewDecoder(resp.Body).Decode(&r)
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
