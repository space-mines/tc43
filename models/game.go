package models

type mine struct {
	x int
	y int
	z int
}

type Sector struct {
	Id        int  `json:"id"`
	X         int  `json:"x"`
	Y         int  `json:"y"`
	Z         int  `json:"z"`
	Radiation int  `json:"radiation"`
	Marked    bool `json:"marked"`
}

type Game struct {
	Id      string   `json:"id"`
	Sectors []Sector `json:"sectors"`
	State   string   `json:"state"`
	mines   []mine
}

var games = make(map[string]Game)

func FindGameById(id string) Game {
	game, exists := games[id]
	if !exists {
		game = CreateNewGame(id, 3, 3)
	}
	return game
}

func CreateNewGame(id string, mineCount int, scale int) Game {
	game := Game{Id: id, State: "PLAY", Sectors: make([]Sector, scale*scale*scale)}
	nextId := 0
	for x := 0; x < scale; x++ {
		for y := 0; y < scale; y++ {
			for z := 0; z < scale; z++ {
				game.Sectors = append(game.Sectors, Sector{Id: nextId, X: x, Y: y, Z: z, Radiation: -1, Marked: false})
				nextId++
			}
		}
	}
	game.mines = []mine{{x: 1, y: 1, z: 1}}
	games[id] = game
	return game
}
