package middleware

import (
	"fmt"
	"net/http"
)

func AuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		fmt.Println("username", user)
		fmt.Println("password", pass)
		if !ok || !checkUsernameAndPassword(user, pass) {
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized.\n"))
			return
		}
		handler(w, r)
	}

}
func checkUsernameAndPassword(username, password string) bool {
	return username == "admin" && password == "admin"
}
