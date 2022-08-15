package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	uh "github.com/derekpedersen/go-utils/http"
	"github.com/derekpedersen/go-utils/mock"

	"github.com/golang/mock/gomock"
)

func TestNewHealthAPIController(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock.NewMockHealthService(ctrl)

	// Act
	r := uh.NewHealthAPIController(svc)

	// Assert
	if r == nil {
		t.Error("expected controller to be set")
	}
}

func TestGetAliveMessage(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock.NewMockHealthService(ctrl)
	api := uh.NewHealthAPIController(svc)

	type args struct {
		method string
		url    string
		params map[string]string
	}
	tests := []struct {
		name       string
		args       args
		statuscode int
	}{
		{
			name: "GetAliveMessage: 200",
			args: args{
				method: http.MethodGet,
				url:    "/alive",
			},
			statuscode: 200,
		},
	}

	msg := uh.HealthMessage{
		Message: "Howdy",
	}

	gomock.InOrder(
		// GetAliveMessage: 200
		svc.EXPECT().GetAliveMessage().Return(&msg),
	)

	// Act
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.args.method, tt.args.url, nil)
			if err != nil {
				t.Fatalf("encountered an unexpected error: %v", err)
			}

			rr := httptest.NewRecorder()
			req = mux.SetURLVars(req, tt.args.params)
			handler := http.HandlerFunc(api.GetAliveMessage)
			handler.ServeHTTP(rr, req)

			// Assert
			if rr.Code != tt.statuscode {
				t.Errorf("unexpected status code: have %v, want %v", rr.Code, tt.statuscode)
			}
		})
	}
}

func TestGetReadyMessage(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock.NewMockHealthService(ctrl)
	api := uh.NewHealthAPIController(svc)

	type args struct {
		method string
		url    string
		params map[string]string
	}
	tests := []struct {
		name       string
		args       args
		statuscode int
	}{
		{
			name: "GetReadyMessage: 200",
			args: args{
				method: http.MethodGet,
				url:    "/ready",
			},
			statuscode: 200,
		},
	}

	msg := uh.HealthMessage{
		Message: "Howdy",
	}

	gomock.InOrder(
		// GetAliveMessage: 200
		svc.EXPECT().GetReadyMessage().Return(&msg),
	)

	// Act
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.args.method, tt.args.url, nil)
			if err != nil {
				t.Fatalf("encountered an unexpected error: %v", err)
			}

			rr := httptest.NewRecorder()
			req = mux.SetURLVars(req, tt.args.params)
			handler := http.HandlerFunc(api.GetReadyMessage)
			handler.ServeHTTP(rr, req)

			// Assert
			if rr.Code != tt.statuscode {
				t.Errorf("unexpected status code: have %v, want %v", rr.Code, tt.statuscode)
			}
		})
	}
}

func TestGetHealthyMessage(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock.NewMockHealthService(ctrl)
	api := uh.NewHealthAPIController(svc)

	type args struct {
		method string
		url    string
		params map[string]string
	}
	tests := []struct {
		name       string
		args       args
		statuscode int
	}{
		{
			name: "GetHealthyMessage: 200",
			args: args{
				method: http.MethodGet,
				url:    "/healthy",
			},
			statuscode: 200,
		},
	}

	msg := uh.HealthMessage{
		Message: "Howdy",
	}

	gomock.InOrder(
		// GetHealthyMessage: 200
		svc.EXPECT().GetHealthyMessage().Return(&msg),
	)

	// Act
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.args.method, tt.args.url, nil)
			if err != nil {
				t.Fatalf("encountered an unexpected error: %v", err)
			}

			rr := httptest.NewRecorder()
			req = mux.SetURLVars(req, tt.args.params)
			handler := http.HandlerFunc(api.GetHealthyMessage)
			handler.ServeHTTP(rr, req)

			// Assert
			if rr.Code != tt.statuscode {
				t.Errorf("unexpected status code: have %v, want %v", rr.Code, tt.statuscode)
			}
		})
	}
}
