package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(New().IndexHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code:got %v want %v", status, http.StatusOK)
	}

	expected := `Welcome to shortlink API`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body : got %v want %v", rr.Body.String(), expected)
	}
}
