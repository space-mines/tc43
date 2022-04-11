package services

import (
	"github.com/heroku/tc43/models"
)

var games = make(map[string]models.Game)

func FindGameById(id string) models.Game {
	game, exists := games[id]
	if !exists {
		game = CreateNewGame(id, 3, 3)
	}
	return game
}

func CreateNewGame(id string, mineCount int, scale int) models.Game {
	game := models.NewGame(id, scale)
	games[id] = game
	return game
}

func MarkSector(id string, sectorId string) models.Game {
	game := FindGameById(id)
	game.Mark(sectorId)
	return game
}

func RevealSector(id string, sectorId string) models.Game {
	game := FindGameById(id)
	game.Reveal(sectorId)
	return game
}
