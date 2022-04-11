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

func NewGame(id string, scale int) Game {
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
	return game
}

func (game *Game) Reveal(sectorId string) {
	sector := game.sectorMap[sectorId]
	sector.Radiation = 1
}

func (game *Game) Mark(sectorId string) {
	sector := game.sectorMap[sectorId]
	println(sector.Id)
	sector.Marked = true
}
