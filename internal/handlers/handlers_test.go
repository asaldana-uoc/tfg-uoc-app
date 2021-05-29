// handlers_test.go
package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Estructura de dades amb tots els endpoints que s'executaran els tests
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

// Funció per a testejar tots els handlers de la classe handlers.go
func TestHandler(t *testing.T) {
	// Iterem sobre tots els elements definits en l'estructura de dades endpointsToCheck
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
