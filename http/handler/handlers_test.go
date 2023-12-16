package handler

import (
	"bytes"
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

type called struct {
	v bool
}

func (s *called) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	s.v = true
}

type dummyWriter struct {
	h  http.Header
	b  bytes.Buffer
	st int
}

func (d *dummyWriter) Header() http.Header {
	if d.h == nil {
		d.h = http.Header{}
	}
	return d.h
}

func (d *dummyWriter) Write(b []byte) (int, error) {
	return d.b.Write(b)
}

func (d *dummyWriter) WriteHeader(statusCode int) {
	d.st = statusCode
}

func TestDenyNoLHttpRequest(t *testing.T) {
	t.Run("go next", func(t *testing.T) {
		c := called{}
		handler := DenyNoLHttpRequest(&c, 200, "")
		header := http.Header{}
		header.Set("X-Requested-With", "XMLHttpRequest")
		handler.ServeHTTP(nil, &http.Request{Header: header})
		if !c.v {
			t.Errorf("should call next handler")
		}
	})
	t.Run("cut off", func(t *testing.T) {
		c := called{}
		h := DenyNoLHttpRequest(&c, http.StatusForbidden, "foo")
		rw := &dummyWriter{}
		h.ServeHTTP(rw, &http.Request{})
		if c.v {
			t.Errorf("should not call next handler")
		}
		if rw.st != http.StatusForbidden {
			t.Errorf("wont=%d, got=%d", http.StatusForbidden, rw.st)
		}
		if got := rw.b.String(); got != "foo" {
			t.Errorf(`wont="foo", got=%q`, got)
		}
	})
}
