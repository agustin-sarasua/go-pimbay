package route

import (
	"encoding/base64"
	"net/http"
	"strings"

	"golang.org/x/net/context"

	"fmt"

	"github.com/agustin-sarasua/pimbay"
	"github.com/agustin-sarasua/pimbay/app/service"
)

const userIdKey = 999

func ValidateToken(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		t := r.Header.Get("Token")
		if t == "" {
			http.Error(w, "Not authorized", 401)
			return
		}

		c := service.GetAccountInfo(t)
		rs := <-c
		if rs == nil {
			http.Error(w, "Not authorized", 401)
			return
		}
		u, _ := pimbay.DB.GetUserByFirebaseID(context.Background(), rs.Users[0].LocalID)

		ctx := context.WithValue(context.Background(), userIdKey, u.ID)
		r = r.WithContext(ctx)
		fmt.Println(u)

		h.ServeHTTP(w, r)
	}
}

func BasicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			http.Error(w, "Not authorized", 401)
			return
		}

		b, err := base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			http.Error(w, "Not authorized", 401)
			return
		}

		c := service.SigninUser(pair[0], pair[1])
		rs := <-c
		if rs == nil && rs.IDToken == "" {
			http.Error(w, "Not authorized", 401)
			return
		}
		w.Header().Set("token", rs.IDToken)
		h.ServeHTTP(w, r)
	}
}
