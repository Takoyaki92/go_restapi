package database

import (
	"fmt"
	"restapi/models"
)

// to do: make mock check what the http layer is sending to the mock (to the db layer)
// depending on that.. if its the value u expect, then return the expected data from the mock.
// if the data is wrong in any way, return an error

type DatabaseMock struct {
	Games          []models.Game
	MockOpen       func() error
	MockGetGames   func() (string, error)
	MockGetGame    func() (string, error)
	MockCreateGame func() error
	MockDeleteGame func() error
	MockUpdateGame func() error
}

func NewMock() DatabaseMock {
	return DatabaseMock{}
}

func (d *DatabaseMock) Open() error {
	return d.MockOpen()
	// return a func
}

// make OpenSuccess () OpenFail, etc.

func OpenSuccess() error {
	return nil
}

func OpenFail() error {
	return fmt.Errorf("failed to initialize DB")
}

func (d *DatabaseMock) Close() error {
	return nil
}

func (d *DatabaseMock) Prepare() error {
	return nil
}

func (d *DatabaseMock) GetGames() (string, error) {
	return d.MockGetGames()
}

func GetGamesSuccess() (string, error) {
	return "Success", nil
}

func GetGamesFail() error {
	return fmt.Errorf("failed to get games")
}

// func (d *DatabaseMock) GetGame(params map[string]string) (models.Game, error) {
func (d *DatabaseMock) GetGame(params map[string]string) (string, error) {
	// id, err := strconv.Atoi(params["id"])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// return d.Games[id], nil
	return d.MockGetGame()
}

func GetGameSuccess() (string, error) {
	return "Game", nil
}

func GetGameFail() error {
	return fmt.Errorf("failed to get games")
}

func (d *DatabaseMock) CreateGame(game models.Game) error {
	return d.MockCreateGame()
}

func CreateGameSuccess() error {
	return nil
}

func CreateGameFail() error {
	return fmt.Errorf("failed to create game")
}

func (d *DatabaseMock) UpdateGame(params map[string]string, game models.Game) error {
	return d.MockUpdateGame()
}

func UpdateGameSuccess() error {
	return nil
}

func UpdateGameFail() error {
	return fmt.Errorf("fauled to update game")
}

func (d *DatabaseMock) DeleteGame(params map[string]string) error {
	return d.MockDeleteGame()
}

func DeleteGameSuccess() error {
	return nil
}

func DeleteGameFail() error {
	return fmt.Errorf("failed to delete game")
}
