package models

type mine struct {
	x int
	y int
	z int
}

type Asteroid struct {
	X         int  `json:"x"`
	Y         int  `json:"y"`
	Z         int  `json:"z"`
	Radiation int  `json:"radiation"`
	Flagged   bool `json:"flagged"`
}

type Game struct {
	Id        string     `json:"id"`
	Asteroids []Asteroid `json:"asteroids"`
	mines     []mine
}

var games = make(map[string]Game)

func FindGameById(id string) Game {
	game, exists := games[id]
	if !exists {
		game = createNewGame(id, 3, 3)
	}
	return game
}

func createNewGame(id string, mineCount int, scale int) Game {
	game := Game{Id: id, Asteroids: make([]Asteroid, scale*scale*scale)}
	for x := 0; x < scale; x++ {
		for y := 0; y < scale; y++ {
			for z := 0; z < scale; z++ {
				game.Asteroids = append(game.Asteroids, Asteroid{X: x, Y: y, Z: z, Radiation: -1, Flagged: false})
			}
		}
	}
	game.mines = []mine{{x: 1, y: 1, z: 1}}
	games[id] = game
	return game
}
