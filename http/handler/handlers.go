package handler

import (
	"net/http"

	"github.com/dairyo/golang-utils/http/request"
	"github.com/dairyo/golang-utils/http/response"
)

// AddNoSniff returns handler to add nosniff as X-Content-Type-Options response header.
func AddNoSniff(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		response.NoSniff(rw)
		next.ServeHTTP(rw, r)
	})
}

// DenyNoLHttpRequest terminates ServeHTTP chain to avoid CSRF.
// If request header does not contains X-Requested-With" header and
// the value is not XMLHttpRequest, HTTP session terminates.
func DenyNoLHttpRequest(next http.Handler, statusCode int, msg string) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if !request.HasXMLHttpRequest(r) {
			rw.WriteHeader(statusCode)
			if msg != "" {
				rw.Write([]byte(msg))
			}
			return
		}
		next.ServeHTTP(rw, r)
	})
}
