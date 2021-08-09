package database

import (
	"restapi/models"
)

type DatabaseMock struct{}

func NewMock() DatabaseMock {
	return DatabaseMock{}
}

func (d *DatabaseMock) Open() error {
	return nil
}

func (d *DatabaseMock) Close() error {
	return nil
}

func (d *DatabaseMock) Prepare() error {
	return nil
}

func (d *DatabaseMock) GetGames() (string, error) {
	return "test", nil
}

// func (d *DatabaseMock) GetGame(params map[string]string) (models.Game, error) {
func (d *DatabaseMock) GetGame(params map[string]string) (string, error) {
	return "test", nil
}

func (d *DatabaseMock) CreateGame(game models.Game) error {
	return nil
}

func (d *DatabaseMock) UpdateGame(params map[string]string, game models.Game) error {
	return nil
}

func (d *DatabaseMock) DeleteGame(params map[string]string) error {
	return nil
}
