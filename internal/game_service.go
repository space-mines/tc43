package internal

var games = make(map[string]Game)

func FindGameById(id string) Game {
	game, exists := games[id]
	if !exists {
		game = CreateNewGame(id, 3, 3)
	}
	return game
}

func CreateNewGame(id string, mineCount int, scale int) Game {
	game := GenerateGame(id, scale)
	games[id] = game
	return game
}

func MarkSector(id string, sectorId string) Game {
	game := FindGameById(id)
	game.Mark(sectorId)
	return game
}

func RevealSector(id string, sectorId string) Game {
	game := FindGameById(id)
	game.Reveal(sectorId)
	return game
}
