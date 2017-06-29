package main_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/agustin-sarasua/pimbay/api"
	"github.com/agustin-sarasua/pimbay/main"
	"github.com/agustin-sarasua/pimbay/model"

	"bytes"

	"github.com/agustin-sarasua/pimbay"
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
)

func TestSignupNewUser(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()
	jsonValue, _ := json.Marshal(api.SignupUserRestMsg{Name: "Agustin", LastName: "Sarasua", Email: "test@test.com", Password: "pwdTest1234"})
	req, err := inst.NewRequest("POST", "/signup", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Failed to create req1: %v", err)
	}
	c1 := appengine.NewContext(req)

	handler := http.HandlerFunc(main.SignupNewUserEndpoint)
	req = req.WithContext(c1)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	q := datastore.NewQuery("User")

	var r []*model.User

	q.GetAll(c1, &r)

	if len(r) != 1 {
		t.Errorf("Mal")
	}

}

func init() {
	fmt.Println("Running init test...")
	pimbay.FbAPI = api.NewFirebaseMockedAPI()
}
