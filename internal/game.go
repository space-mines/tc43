package internal

import (
	"fmt"
	"strconv"
)

type Game struct {
	Id      string   `json:"id"`
	Size    int      `json:"size"`
	Sectors []Sector `json:"sectors"`
	State   string   `json:"state"`
	mines   []Location
}

func (sector Sector) print() {
	println(fmt.Sprintf("Sector[id=%d,radiation=%d,marked=%v]", sector.Id, sector.Radiation, sector.Marked))
}

func NewGame(id string, mines []Location, sectors []Sector, scale int) Game {
	if sectors == nil || len(sectors) == 0 {
		sectors = GenerateBlankSectors(scale)
	}
	return Game{Id: id, State: "PLAY", Sectors: sectors, Size: scale, mines: mines}
}

func GenerateBlankSectors(scale int) []Sector {
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
	sectors := GenerateBlankSectors(scale)
	game := Game{Id: id, State: "PLAY", Sectors: sectors}
	game.mines = []Location{{X: 1, Y: 1, Z: 1}, {X: 0, Y: 0, Z: 0}}
	return game
}

func (game *Game) Reveal(sectorId string) {
	index, _ := strconv.Atoi(sectorId)
	sector := &game.Sectors[index]
	location := Location{sector.X, sector.Y, sector.Z}
	radiation := location.CalculateRadiation(game.mines)
	sector.Radiation = radiation
	sector.print()
}

func (game *Game) Mark(sectorId string) {
	index, _ := strconv.Atoi(sectorId)
	sector := &game.Sectors[index]
	sector.Marked = true
	sector.print()
}
