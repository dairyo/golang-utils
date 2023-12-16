package handler

import (
	"net/http"

	"github.com/dairyo/golang-utils/http/response"
)

// AddNoSniff returns handler to add nosniff as X-Content-Type-Options response header.
func AddNoSniff(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		response.NoSniff(rw)
		next.ServeHTTP(rw, r)
	})
}
