package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Game struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Developer string `json:"developer"`
	Rating    string `json:"rating"`
}

func getGames(w http.ResponseWriter, r *http.Request) {
	database, err := sql.Open("sqlite3", "./games.db")
	if err != nil {
		fmt.Println(err)
	}

	rows, err := database.Query("SELECT * FROM games")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var games []*Game

	for rows.Next() {
		g := new(Game)
		err := rows.Scan(&g.ID, &g.Title, &g.Developer, &g.Rating)
		if err != nil {
			fmt.Println(err)
		}

		games = append(games, g)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func getGame(w http.ResponseWriter, r *http.Request) {
	database, err := sql.Open("sqlite3", "./games.db")
	if err != nil {
		fmt.Println(err)
	}

	params := mux.Vars(r)

	rows, err := database.Query("SELECT * FROM games WHERE id=?", params["id"])
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		g := new(Game)
		err := rows.Scan(&g.ID, &g.Title, &g.Developer, &g.Rating)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(g)
	}

}

func createGame(w http.ResponseWriter, r *http.Request) {
	database, err := sql.Open("sqlite3", "./games.db")
	if err != nil {
		fmt.Println(err)
	}

	var game Game

	_ = json.NewDecoder(r.Body).Decode(&game)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
	database.Exec("INSERT INTO games VALUES(?, ?, ?, ?)", game.ID, game.Title, game.Developer, game.Rating)

}

func updateGame(w http.ResponseWriter, r *http.Request) {
	database, err := sql.Open("sqlite3", "./games.db")
	if err != nil {
		fmt.Println(err)
	}

	params := mux.Vars(r)

	database.Exec("DELETE FROM games WHERE id=?", params["id"])

	var game Game

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&game)
	json.NewEncoder(w).Encode(game)

	database.Exec("INSERT INTO games VALUES(?, ?, ?, ?)", game.ID, game.Title, game.Developer, game.Rating)
}

func deleteGame(w http.ResponseWriter, r *http.Request) {
	database, err := sql.Open("sqlite3", "./games.db")
	if err != nil {
		fmt.Println(err)
	}

	params := mux.Vars(r)

	database.Exec("DELETE FROM games WHERE id=?", params["id"])
}

func main() {
	database, err := sql.Open("sqlite3", "./games.db")
	if err != nil {
		fmt.Println(err)
	}

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS games (id INTEGER PRIMARY KEY, title TEXT, developer TEXT, rating TEXT)")
	statement.Exec()

	defer database.Close()

	// Init Router
	r := mux.NewRouter()

	// Route handlers / endpoints
	r.HandleFunc("/api/games", getGames).Methods("GET")
	r.HandleFunc("/api/games/{id}", getGame).Methods("GET")
	r.HandleFunc("/api/games", createGame).Methods("POST")
	r.HandleFunc("/api/games/{id}", updateGame).Methods("PUT")
	r.HandleFunc("/api/games/{id}", deleteGame).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
