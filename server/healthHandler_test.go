package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	srv := New(nil, "1.0.0")

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(srv.healthCheck())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("HTTP Error: Expected: %v | Got: %v", http.StatusOK, status)
	}

	expected := `{"status":"ok"}`
	body := strings.TrimSpace(rr.Body.String())
	if body != expected {
		t.Errorf("unexpected body: Expected: %v | Got: %v", expected, body)
	}
}

func TestVersionCheck(t *testing.T) {
	srv := New(nil, "1.0.0")

	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(srv.versionCheck())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("HTTP Error: Expected: %v | Got: %v", http.StatusOK, status)
	}

	expected := `{"version":"1.0.0"}`
	body := strings.TrimSpace(rr.Body.String())
	if body != expected {
		t.Errorf("unexpected body: Expected: %v | Got: %v", expected, body)
	}
}
