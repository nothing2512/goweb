package middlewares

import (
	"net/http"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Do Auth Here
		next.ServeHTTP(w, r)
	}
}
