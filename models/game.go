package models

import "strconv"

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
	Id        string   `json:"id"`
	Sectors   []Sector `json:"sectors"`
	State     string   `json:"state"`
	sectorMap map[string]Sector
	mines     []mine
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
	game := Game{Id: id, State: "PLAY", Sectors: make([]Sector, 0), sectorMap: make(map[string]Sector)}
	nextId := 0
	for x := 0; x < scale; x++ {
		for y := 0; y < scale; y++ {
			for z := 0; z < scale; z++ {
				sector := Sector{Id: nextId, X: x, Y: y, Z: z, Radiation: -1, Marked: false}
				game.Sectors = append(game.Sectors, sector)
				game.sectorMap[strconv.Itoa(nextId)] = sector
				nextId++
			}
		}
	}
	game.mines = []mine{{x: 1, y: 1, z: 1}}
	games[id] = game
	return game
}

func RevealSector(id string, sectorId string) Game {
	game := FindGameById(id)
	sector := game.sectorMap[sectorId]
	sector.Radiation = 1
	return game
}

func MarkSector(id string, sectorId string) Game {
	game := FindGameById(id)
	sector := game.sectorMap[sectorId]
	println(sector.Id)
	sector.Marked = true
	return game
}
