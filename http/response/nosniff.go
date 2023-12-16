package response

import "net/http"

// NoSniff adds X-Content-Type-Options header to a response.
func NoSniff(w http.ResponseWriter) {
	w.Header().Add("X-Content-Type-Options", "nosniff")
}
