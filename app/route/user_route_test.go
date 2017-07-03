package route_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"cloud.google.com/go/datastore"

	"golang.org/x/net/context"

	"github.com/agustin-sarasua/pimbay"
	"github.com/agustin-sarasua/pimbay/app/api"
	"github.com/agustin-sarasua/pimbay/app/db"
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

func init() {
	fmt.Println("Running init test...")
	os.Setenv("DATASTORE_DATASET", "pimbay-accounting")
	os.Setenv("DATASTORE_EMULATOR_HOST", "localhost:8081")
	os.Setenv("DATASTORE_EMULATOR_HOST_PATH", "localhost:8081/datastore")
	os.Setenv("DATASTORE_HOST", "http://localhost:8081")
	os.Setenv("DATASTORE_PROJECT_ID", "pimbay-accounting")
	pimbay.FbAPI = api.NewFirebaseMockedAPI()
	pimbay.DB, _ = configureDatastoreDB("pimbay-accounting")
}

func configureDatastoreDB(projectID string) (db.Database, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return db.NewDatastoreDB(ctx, client)
}
