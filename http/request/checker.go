package request

import "net/http"

// HasXMLHttpRequest checks whether request header contains XMLHttpRequest as value of X-Requested-With.
func HasXMLHttpRequest(r *http.Request) bool {
	return CheckRequestedWith(r, "XMLHttpRequest")
}

// CheckRequestedWith checks X-Requested-With value.
// It can be use Prevention.
func CheckRequestedWith(r *http.Request, expected string) bool {
	actual := r.Header.Get("X-Requested-With")
	if actual != expected {
		return false
	}
	return true
}
