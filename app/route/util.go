package route

import (
	"fmt"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func GetUserID(req *http.Request) int64 {
	ctx := req.Context()
	uid := ctx.Value(999)
	fmt.Println("User Id: ", uid)
	fmt.Println("Hello")
	return uid.(int64)
}
