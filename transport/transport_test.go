package transport_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"restapi/transport"
	"testing"
)

func TestGetGames(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/games", nil)
	if err != nil {
		log.Fatal(err)
	}

	// rr is ResponseRecorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(transport.GetGames)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, wanted %v", status, http.StatusOK)
	}

	expected := `[{"id":"1","title":"Sonic Unleashed","developer":"Sonic Team","rating":"5"},{"id":"2","title":"Sonic 2","developer":"Sega","rating":"10"},{"id":"3","title":"Sonic 3","developer":"Sega","rating":"10"}]`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v, wanted %v", rr.Body.String(), expected)
	}
}

func TestGetGame(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/games", nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(transport.GetGame)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"1","title":"Sonic Unleashed","developer":"Sonic Team","rating":"5"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
