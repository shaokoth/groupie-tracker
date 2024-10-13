package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
)

var mockArtistAPI = `{"ArtistID": 1,
 					"ArtistName": "Test Artist",
					"ArtistImage": "Image url",
					"BandMembers": ["Member1", "Member2", "Member3"],
					"CreationDate": "1/1/2024",
					"FirstAlbum": "1/2/2024"}`

// Helper function to create a mock server for dates API
func MockerServer(mockResponse string) *httptest.Server {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, mockResponse)
	}))
	return mockServer
}

