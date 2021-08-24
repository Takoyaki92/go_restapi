package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"restapi/database"
	"restapi/models"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// var DB = database.New()
var DB = database.DatabaseMock{MockOpen: database.OpenSuccess, MockGetGames: database.GetGamesSuccess, MockGetGame: database.GetGameSuccess, MockCreateGame: database.CreateGameSuccess, MockDeleteGame: database.DeleteGameSuccess, MockUpdateGame: database.UpdateGameSuccess}

// to set the dummy DB for integration tests, change var dbPath in database.go as well

func GetGames(w http.ResponseWriter, r *http.Request) {
	// games, err := DB.GetGames()
	games, err := DB.GetGames()
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func GetGame(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	g, err := DB.GetGame(params)
	if err != nil {
		fmt.Println(err)
	}

	// defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(g)

}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	var game models.Game

	_ = json.NewDecoder(r.Body).Decode(&game)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)

	err := DB.CreateGame(game)
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateGame(w http.ResponseWriter, r *http.Request) {
	var game models.Game

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&game)
	json.NewEncoder(w).Encode(game)

	params := mux.Vars(r)
	err := DB.UpdateGame(params, game)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := DB.DeleteGame(params)
	if err != nil {
		fmt.Println(err)
	}
}
