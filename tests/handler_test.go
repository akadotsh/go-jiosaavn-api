package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/akadotsh/go-jiosaavn-client/api"
)
func TestAPIRootHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/api/", nil)
    if err != nil {
        t.Fatal(err)
    }

    recorder := httptest.NewRecorder()

    handler := http.HandlerFunc(api.RootHandler)

    handler.ServeHTTP(recorder, req)

    if status := recorder.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    contentType := recorder.Header().Get("Content-Type")
    expectedContentType := "text/html; charset=utf-8"
    if contentType != expectedContentType {
        t.Errorf("handler returned wrong content type: got %v want %v",
            contentType, expectedContentType)
    }

    body := recorder.Body.String()
    expectedContents := []string{
        "API Routes",
        "/api/songs",
        "/api/artists",
        "/api/playlists",
        "/api/search",
    }

    for _, content := range expectedContents {
        if !strings.Contains(body, content) {
            t.Errorf("handler response doesn't contain expected content: %s", content)
        }
    }
}