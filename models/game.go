package models

type asteroid struct {
	Id        string `json:"id"`
	Radiation int    `json:"radiation"`
	Flagged   bool   `json:"flagged"`
}

type game struct {
	Id        string              `json:"id"`
	Asteroids map[string]asteroid `json:"asteroids"`
	mines     []string
}

var games = make(map[string]game)

func findGameById(id string) game {

	return games[id]
}
