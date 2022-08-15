package http_test

import (
	"net/http"
	"testing"

	uh "github.com/derekpedersen/go-utils/http"
)

func TestHttpRequest(t *testing.T) {
	// Arrange
	url := "https://google.com"

	// Act
	response, err := uh.HttpRequest(url, http.MethodGet, nil)

	// Assert
	if err != nil {
		t.Fatalf("An unexpected error: %v", err)
	}
	if response == nil {
		t.Fatalf("No response was retured!")
	}
	if len(*response) == 0 {
		t.Errorf("Expected it to contain a response")
	}
}
