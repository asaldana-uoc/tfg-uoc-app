// handlers_test.go
package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var endpointsToCheck = []struct {
	handlerName      http.HandlerFunc
	URL              string
	method           string
	responseCode     int
	expectedResponse string
}{
	{Home, "/", "GET", http.StatusOK, ""},
	{Status, "/status", "GET", http.StatusOK, "Health check OK"},
	{Home, "/testing", "GET", http.StatusNotFound, "Pàgina no trobada\n"},
}

func TestHandler(t *testing.T) {
	for _, key := range endpointsToCheck {
		req, err := http.NewRequest(key.method, key.URL, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		key.handlerName.ServeHTTP(rr, req)

		if status := rr.Code; status != key.responseCode {
			t.Errorf("Codi HTTP obtingut erroni, s'esperava %v però s'ha rebut  %v",
				status, key.responseCode)
		}

		if key.expectedResponse != "" && rr.Body.String() != key.expectedResponse {
			t.Errorf("Resposta obtinguda invàlida, s'esperava %v però s'ha rebut  %v",
				rr.Body.String(), key.expectedResponse)
		}
	}
}
