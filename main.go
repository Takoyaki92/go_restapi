package main

import (
	"fmt"
	"log"
	"net/http"

	"restapi/transport"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := transport.DB
	err := db.Prepare()
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	// Init Router
	r := mux.NewRouter()

	// Route handlers / endpoints
	r.HandleFunc("/api/games", transport.GetGames).Methods("GET")
	r.HandleFunc("/api/games/{id}", transport.GetGame).Methods("GET")
	r.HandleFunc("/api/games", transport.CreateGame).Methods("POST")
	r.HandleFunc("/api/games/{id}", transport.UpdateGame).Methods("PUT")
	r.HandleFunc("/api/games/{id}", transport.DeleteGame).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
