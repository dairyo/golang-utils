package response

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

func TestNoSniff(t *testing.T) {
	rw := &nosniffCheck{h: http.Header{}}
	NoSniff(rw)
	if got := rw.h.Get("X-Content-Type-Options"); got != "nosniff" {
		t.Errorf("wont = nosniff, got:=%s", got)
	}
}
