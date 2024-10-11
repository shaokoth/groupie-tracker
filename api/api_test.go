package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchAPI(t *testing.T) {
	// Set up a mock server
	mockResponse := `{"message": "hello world"}`
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, mockResponse)
	}))
	defer mockServer.Close()

	// Call the FetchAPI function with the mock server's URL
	result, err := FetchAPI(mockServer.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify the result
	expected := mockResponse
	if string(result) != expected {
		t.Errorf("expected %v, got %v", expected, string(result))
	}
}
