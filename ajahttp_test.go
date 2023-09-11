package ajahttp

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"
)

func TestExtractQueryInUrl(t *testing.T) {
	var rev string
	var err error

	query := url.Values{}
	query.Add("page", "1")
	query.Add("name", "1")

	rev, err = extractQueryInUrl("/user?path=1&name=2", &query)
	if err != nil {
		t.Fatal()
	}

	if rev != "/user" {
		t.Errorf("Reverse: %s, want %s", rev, "/user")
	}

	if !query.Has("path") {
		t.Errorf("Reverse: %v, want %v", false, true)
	}

	if !query.Has("name") {
		t.Errorf("Reverse: %v, want %v", false, true)
	}

	if !query.Has("page") {
		t.Errorf("Reverse: %v, want %v", false, true)
	}

	if rev := query.Get("page"); rev != "1" {
		t.Errorf("Reverse: %v, want %v", rev, "1")
	}

	if rev := query.Get("path"); rev != "1" {
		t.Errorf("Reverse: %v, want %v", rev, "1")
	}

	rev2 := query["name"]
	if rev2[0] != "1" || rev2[1] != "2" {
		t.Fatal()
	}
}

func TestGetBody(t *testing.T) {
	// $ nc -lvp 3344

	opt := &AjaOption{
		Method:  http.MethodGet,
		Params:  url.Values{"page": {"1"}, "name": {"foo", "bar"}},
		Url:     "http://localhost:3344?page=2",
		Headers: map[string]string{"X-My": "test"},
		Body:    bytes.NewBufferString("abc"),
	}

	_, err := Request(opt)

	if err != nil {
		t.Fatal()
	}
}
