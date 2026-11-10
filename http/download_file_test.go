package http_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	uh "github.com/derekpedersen/go-utils/http"
)

func TestDownloadToFile(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Test") != "yes" {
			t.Fatalf("missing request header")
		}
		_, _ = w.Write([]byte("payload"))
	}))
	defer server.Close()

	destination := filepath.Join(t.TempDir(), "nested", "download.txt")
	if err := uh.DownloadToFile(server.URL, destination, map[string][]string{"X-Test": {"yes"}}); err != nil {
		t.Fatal(err)
	}
	contents, err := os.ReadFile(destination)
	if err != nil {
		t.Fatal(err)
	}
	if string(contents) != "payload" {
		t.Fatalf("got %q", contents)
	}
}

func TestDownloadToFileReturnsResponseError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "failed", http.StatusBadGateway)
	}))
	defer server.Close()

	err := uh.DownloadToFile(server.URL, filepath.Join(t.TempDir(), "download.txt"), nil)
	var responseErr *uh.ResponseError
	if !errors.As(err, &responseErr) {
		t.Fatalf("expected response error, got %T", err)
	}
	if responseErr.StatusCode != http.StatusBadGateway {
		t.Fatalf("got %d", responseErr.StatusCode)
	}
}

func TestDownloadToFileWithClientContextCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := uh.DownloadToFileWithClient(ctx, http.DefaultClient, "http://127.0.0.1:1", filepath.Join(t.TempDir(), "download.txt"), nil)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("got %v", err)
	}
}

func TestDefaultDownloadPath(t *testing.T) {
	path := uh.DefaultDownloadPath("./bin/downloads", "orders")
	if filepath.Base(path) == "" {
		t.Fatal("expected a filename")
	}
	if filepath.Ext(path) != ".json" {
		t.Fatalf("expected .json file, got %s", path)
	}
}
