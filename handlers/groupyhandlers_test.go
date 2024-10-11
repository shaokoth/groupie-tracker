package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

// Test groupyHandler
func TestGroupyHandler(t *testing.T) {
	// Test case: Artist found and template executed successfully
	t.Run("Artist found", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/artist/test", nil)
		rr := httptest.NewRecorder()

		// Use a simple template for testing
		tpl := "artist.tmpl"
		tmplContent := "{{.ArtistName}}"
		template.Must(template.New("artist.tmpl").Parse(tmplContent))

		// Call the handler
		groupyHandler(tpl, rr, req)

		// Check status code
		if rr.Code != http.StatusOK {
			t.Errorf("Expected status code 200, got %v", rr.Code)
		}

		// Check response body
		expectedBody := "Test Artist"
		if rr.Body.String() != expectedBody {
			t.Errorf("Expected body %q, got %q", expectedBody, rr.Body.String())
		}
	})

	// Test case: Artist not found (404)
	t.Run("Artist not found", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/artist/notfound", nil)
		rr := httptest.NewRecorder()

		tpl := "artist.tmpl" // Template won't matter for 404

		groupyHandler(tpl, rr, req)

		// Check status code
		if rr.Code != http.StatusNotFound {
			t.Errorf("Expected status code 404, got %v", rr.Code)
		}

		// Check response body for error message
		expectedErr := "Paths cross empty void\nSeeking what once existed\nSilence answers all."
		if rr.Body.String() != expectedErr {
			t.Errorf("Expected error message %q, got %q", expectedErr, rr.Body.String())
		}
	})

	// Test case: Internal server error (template parsing error)
	t.Run("Internal server error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/artist/test", nil)
		rr := httptest.NewRecorder()

		// Use an invalid template path to simulate template parsing error
		tpl := "invalid-template.tmpl"

		groupyHandler(tpl, rr, req)

		// Check status code
		if rr.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code 500, got %v", rr.Code)
		}

		// Check response body for internal server error message
		expectedErr := "Internal Server Error"
		if rr.Body.String() != expectedErr {
			t.Errorf("Expected error message %q, got %q", expectedErr, rr.Body.String())
		}
	})
}
