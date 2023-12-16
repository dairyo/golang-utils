package handler

import (
	"net/http"
	"testing"
)

type nosniffCheck struct {
	t *testing.T
	h http.Header
}

func (n *nosniffCheck) Header() http.Header {
	return n.h
}

func (n *nosniffCheck) Write([]byte) (int, error) {
	n.t.Helper()
	n.t.Fatal("Write called.")
	return 0, nil
}

func (n *nosniffCheck) WriteHeader(statusCode int) {
	n.t.Helper()
	n.t.Fatal("WriteHeader called.")
}

func TestAddNoSniff(t *testing.T) {
	check := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		t.Helper()
		if got := rw.Header().Get("X-Content-Type-Options"); got != "nosniff" {
			t.Errorf("wont=sniff, got=%s", got)
		}
	})
	h := AddNoSniff(check)
	h.ServeHTTP(&nosniffCheck{h: http.Header{}}, &http.Request{})
}
