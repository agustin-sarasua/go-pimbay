package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/agustin-sarasua/pimbay/api"
	"github.com/agustin-sarasua/pimbay/main"

	"bytes"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

func TestSignupNewUser(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()
	jsonValue, _ := json.Marshal(api.SignupUserRestMsg{Name: "Agustin", LastName: "Sarasua", Email: "test2@test.com", Password: "pwdTest1234"})
	req, err := inst.NewRequest("POST", "/signup", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Failed to create req1: %v", err)
	}
	c1 := appengine.NewContext(req)

	handler := http.HandlerFunc(main.SignupNewUserEndpoint)
	req = req.WithContext(c1)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// c1 := appengine.NewContext(req1)

	// req2, err := inst.NewRequest("GET", "/herons", nil)
	// if err != nil {
	// 	t.Fatalf("Failed to create req2: %v", err)
	// }
	// c2 := appengine.NewContext(req2)

}
