package ajahttp

import (
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
