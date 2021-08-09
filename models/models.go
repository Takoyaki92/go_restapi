package models

type Game struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Developer string `json:"developer"`
	Rating    string `json:"rating"`
}
