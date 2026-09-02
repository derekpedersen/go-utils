package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	uh "github.com/derekpedersen/go-utils/http"
)

type healthService struct {
	message *uh.HealthMessage
}

func (service healthService) GetAliveMessage() *uh.HealthMessage   { return service.message }
func (service healthService) GetReadyMessage() *uh.HealthMessage   { return service.message }
func (service healthService) GetHealthyMessage() *uh.HealthMessage { return service.message }

func TestHealthAPIController(t *testing.T) {
	api := uh.NewHealthAPIController(healthService{message: &uh.HealthMessage{Message: "Howdy"}})
	tests := []struct {
		name    string
		handler http.HandlerFunc
	}{
		{name: "alive", handler: api.GetAliveMessage},
		{name: "ready", handler: api.GetReadyMessage},
		{name: "healthy", handler: api.GetHealthyMessage},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			test.handler.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/", nil))

			if recorder.Code != http.StatusOK {
				t.Fatalf("got status %d", recorder.Code)
			}
			if recorder.Header().Get("Content-Type") != "application/json" {
				t.Fatalf("got content type %q", recorder.Header().Get("Content-Type"))
			}
			var message uh.HealthMessage
			if err := json.Unmarshal(recorder.Body.Bytes(), &message); err != nil {
				t.Fatal(err)
			}
			if message.Message != "Howdy" {
				t.Fatalf("got message %q", message.Message)
			}
		})
	}
}
