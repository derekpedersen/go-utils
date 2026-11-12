package http_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestHttpRequestWithBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(payload) != "{\"name\":\"Ada\"}" {
			t.Fatalf("unexpected payload: %q", payload)
		}
		_, _ = w.Write([]byte("created"))
	}))
	defer server.Close()

	response, err := uh.HttpRequestWithBody(server.URL, http.MethodPost, map[string][]string{"Content-Type": {"application/json"}}, strings.NewReader("{\"name\":\"Ada\"}"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response == nil || *response != "created" {
		t.Fatalf("unexpected response: %v", response)
	}
}

func TestHttpRequestWithContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := uh.HttpRequestWithContext(ctx, "http://127.0.0.1:1", http.MethodGet, nil)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("got %v, want context canceled", err)
	}
}

func TestIsSuccessStatus(t *testing.T) {
	if !uh.IsSuccessStatus(http.StatusNoContent) {
		t.Fatal("expected 204 to be successful")
	}
	if uh.IsSuccessStatus(http.StatusBadGateway) {
		t.Fatal("expected 502 to be non-successful")
	}
}
