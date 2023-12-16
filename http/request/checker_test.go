package request

import (
	"net/http"
	"testing"
)

func TestHasXMLHttpRequest(t *testing.T) {
	genReq := func(v string) *http.Request {
		h := http.Header{}
		h.Add("X-Requested-With", v)
		return &http.Request{Header: h}
	}
	check := func(t *testing.T, r *http.Request, wont bool) {
		t.Helper()
		got := HasXMLHttpRequest(r)
		if wont != got {
			t.Errorf("wont: %t, got: %t", wont, got)
		}
	}
	check(t, genReq("XMLHttpRequest"), true)
	check(t, &http.Request{}, false)
	check(t, genReq("foo"), false)
}

func TestCheckRequestedWith(t *testing.T) {
	check := func(t *testing.T, inHeader, expected string, wont bool) {
		t.Helper()
		h := http.Header{}
		h.Add("X-Requested-With", inHeader)
		r := &http.Request{Header: h}
		got := CheckRequestedWith(r, expected)
		if wont != got {
			t.Errorf("wont: %t, got: %t", wont, got)
		}
	}
	check(t, "foo", "foo", true)
	check(t, "foo", "bar", false)
	check(t, "foo", "", false)
}
