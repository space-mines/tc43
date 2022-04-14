package internal

import (
	"fmt"
	"strconv"
)

type Location struct {
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
	mines   []Location
}

func (sector Sector) print() {
	println(fmt.Sprintf("Sector[id=%d,marked=%v]", sector.Id, sector.Marked))
}

func NewGame(id string, mines []Location, sectors []Sector, scale int) Game {
	if sectors == nil || len(sectors) == 0 {
		sectors = generateBlankSectors(scale)
	}
	return Game{Id: id, State: "PLAY", Sectors: sectors, mines: mines}
}

func generateBlankSectors(scale int) []Sector {
	nextId := 0
	sectors := make([]Sector, 0)
	for x := 0; x < scale; x++ {
		for y := 0; y < scale; y++ {
			for z := 0; z < scale; z++ {
				sector := Sector{Id: nextId, X: x, Y: y, Z: z, Radiation: -1, Marked: false}
				sectors = append(sectors, sector)
				nextId++
			}
		}
	}
	return sectors
}

func GenerateGame(id string, scale int) Game {
	sectors := generateBlankSectors(scale)
	game := Game{Id: id, State: "PLAY", Sectors: sectors}
	game.mines = []Location{{x: 1, y: 1, z: 1}}
	return game
}

func (game *Game) Reveal(sectorId string) {
	index, _ := strconv.Atoi(sectorId)
	sector := &game.Sectors[index]
	sector.Radiation = 1
	sector.print()
}

func (game *Game) Mark(sectorId string) {
	index, _ := strconv.Atoi(sectorId)
	sector := &game.Sectors[index]
	sector.Marked = true
	sector.print()
}
