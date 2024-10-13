package tests

import (
	"testing"

	"groupie-tracker/api"
)

func TestFetchAPI(t *testing.T) {
	// mockResponse := `{"message": "hello world"}`
	server := MockerServer(mockArtistAPI)

	defer server.Close()

	result, err := api.FetchAPI(server.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	// Verify the result
	expected := mockArtistAPI
	if string(result) != expected {
		t.Errorf("expected %v, got %v", expected, string(result))
	}
}
