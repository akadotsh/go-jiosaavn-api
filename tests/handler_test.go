package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akadotsh/go-jiosaavn-client/api"
	"github.com/akadotsh/go-jiosaavn-client/utils"
)

func TestRootHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(api.RootHandler)

	handler.ServeHTTP(recorder, req)

	fmt.Println("recorder", recorder)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var resp utils.Response
	json.Unmarshal(recorder.Body.Bytes(), &resp)

	expectedStatus := "success"
	expectedMessage := "Beep Boop!"

	if resp.Status != expectedStatus {
		t.Errorf("handler returned unexpected status: got %v want %v",
			resp.Status, expectedStatus)
	}

	if resp.Message != expectedMessage {
		t.Errorf("handler returned unexpected message: got %v want %v",
			resp.Message, expectedMessage)
	}

}
