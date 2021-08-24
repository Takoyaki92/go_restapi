package transport_test

import (
	"bytes"
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

	expected := "\"Success\"\n"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got [%v], wanted [%v]", rr.Body.String(), expected)
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
	expected := "\"Game\"\n"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreateGame(t *testing.T) {
	// create new HTTP post request
	var json = []byte(`{"title": "title", "developer": "developer", "rating": "rating"}`)
	req, _ := http.NewRequest("POST", "/api/games", bytes.NewBuffer(json))
	// assign HTTP handler
	handler := http.HandlerFunc(transport.CreateGame)
	// record response
	rr := httptest.NewRecorder()
	// dispatch http request
	handler.ServeHTTP(rr, req)
	// add assertions
	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"","title":"title","developer":"developer","rating":"rating"}
`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestDeleteGame(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/api/games", nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(transport.DeleteGame)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := ""
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUpdateGame(t *testing.T) {
	// create new HTTP post request
	var json = []byte(`{"title": "title", "developer": "developer", "rating": "rating"}`)
	req, _ := http.NewRequest("PUT", "/api/games", bytes.NewBuffer(json))
	// assign HTTP handler
	handler := http.HandlerFunc(transport.UpdateGame)
	// record response
	rr := httptest.NewRecorder()
	// dispatch http request
	handler.ServeHTTP(rr, req)
	// add assertions
	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"","title":"title","developer":"developer","rating":"rating"}
`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}
