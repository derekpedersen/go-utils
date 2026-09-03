package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	uh "github.com/derekpedersen/go-utils/http"
)

func TestHttpRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Test") != "yes" {
			t.Errorf("missing request header")
		}
		_, _ = w.Write([]byte("ok"))
	}))
	defer server.Close()

	response, err := uh.HttpRequest(server.URL, http.MethodGet, map[string][]string{"X-Test": {"yes"}})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response == nil {
		t.Fatal("no response returned")
	}
	if *response != "ok" {
		t.Errorf("got %q, want %q", *response, "ok")
	}
}

func TestHttpRequestReturnsResponseErrorForNon2xx(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "nope", http.StatusBadRequest)
	}))
	defer server.Close()

	response, err := uh.HttpRequest(server.URL, http.MethodGet, nil)
	if response == nil || *response != "nope\n" {
		t.Fatalf("unexpected response: %v", response)
	}
	statusErr, ok := err.(*uh.ResponseError)
	if !ok || statusErr.StatusCode != http.StatusBadRequest {
		t.Fatalf("got %T %v, want status error", err, err)
	}
}
