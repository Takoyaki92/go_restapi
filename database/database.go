package database

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
	"restapi/models"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseI interface {
	Open() error
	Close() error
	Prepare() error
	GetGames() ([]models.Game, error)
	GetGame() (models.Game, error)
	CreateGame() error
	UpdateGame() error
	DeleteGame() error
}

type Database struct {
	db *sql.DB
}

func New() Database {
	return Database{}
}

// var dbPath, _ = filepath.Abs("./database/games.db")
var dbPath, _ = filepath.Abs("./database/test.db")

// Opens the DB
func (d *Database) Open() error {
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	d.db = database
	return nil
}

// Closes the DB
func (d *Database) Close() error {
	return d.db.Close()
}

// Creates the table in the DB if it does not already exist
func (d *Database) Prepare() error {
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS games (id INTEGER PRIMARY KEY, title TEXT, developer TEXT, rating TEXT)")
	statement.Exec()
	return nil
}

// Gets all games from the DB
func (d *Database) GetGames() ([]models.Game, error) {
	var games []models.Game

	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := database.Query("SELECT * FROM games")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var g models.Game
		err := rows.Scan(&g.ID, &g.Title, &g.Developer, &g.Rating)
		if err != nil {
			fmt.Println(err)
		}

		games = append(games, g)
	}
	return games, nil
}

// Gets a single game via ID
func (d *Database) GetGame(params map[string]string) (models.Game, error) {
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := database.Query("SELECT * FROM games WHERE id=?", params["id"])
	if err != nil {
		log.Fatal(err)
	}

	var g models.Game

	for rows.Next() {
		err := rows.Scan(&g.ID, &g.Title, &g.Developer, &g.Rating)
		if err != nil {
			log.Fatal(err)
		}
	}
	return g, nil
}

// Creates a new game in the DB
func (d *Database) CreateGame(game models.Game) error {
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.Exec("INSERT INTO games VALUES(?, ?, ?, ?)", game.ID, game.Title, game.Developer, game.Rating)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// Updates an existing game in the DB
func (d *Database) UpdateGame(params map[string]string, game models.Game) error {
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.Exec("UPDATE games SET title=?, developer=?, rating=? WHERE id=?", game.Title, game.Developer, game.Rating, params["id"])
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// Deletes an existing game in the DB
func (d *Database) DeleteGame(params map[string]string) error {
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.Exec("DELETE FROM games WHERE id=?", params["id"])
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
