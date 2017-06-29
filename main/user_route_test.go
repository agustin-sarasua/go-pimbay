package main_test

import (
	"testing"

	"google.golang.org/appengine/aetest"
)

func TestSignupNewUser(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	_, err = inst.NewRequest("POST", "/signup", nil)
	if err != nil {
		t.Fatalf("Failed to create req1: %v", err)
	}
	// c1 := appengine.NewContext(req)
	// client, err := datastore.NewClient(c1, "pimbay-accounting")
	// b, _ := db.NewDatastoreDB(c1, client)

	// handler := http.HandlerFunc(main.SignupNewUserEndpoint)
	// req = req.WithContext(c1)
	// rr := httptest.NewRecorder()
	// handler.ServeHTTP(rr, req)

	// c1 := appengine.NewContext(req1)

	// req2, err := inst.NewRequest("GET", "/herons", nil)
	// if err != nil {
	// 	t.Fatalf("Failed to create req2: %v", err)
	// }
	// c2 := appengine.NewContext(req2)

}
