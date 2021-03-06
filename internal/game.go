package internal

import (
	"fmt"
	"math"
)

type Game struct {
	Id      string   `json:"id"`
	Size    int      `json:"size"`
	Sectors []Sector `json:"sectors"`
	State   string   `json:"state"`
	mines   []Location
}

func (sector Sector) print() {
	println(fmt.Sprintf("Sector[id=%d,x=%d,y=%d,z=%d,radiation=%d,marked=%v]",
		sector.Id, sector.X, sector.Y, sector.Z, sector.Radiation, sector.Marked))
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
	game := Game{Id: id, State: "PLAY", Sectors: sectors, Size: scale}
	game.mines = []Location{{X: 1, Y: 1, Z: 1}, {X: 0, Y: 0, Z: 0}}
	return game
}

func (game *Game) Reveal(sectorId int) {
	if game.State == "LOSE" || sectorId < 0 || sectorId >= len(game.Sectors) {
		return
	}
	sector := &game.Sectors[sectorId]
	if sector.Radiation != -1 || sector.Marked {
		return
	}
	location := Location{sector.X, sector.Y, sector.Z}
	radiation := location.CalculateRadiation(game.mines)
	sector.Radiation = radiation
	sector.print()
	if sector.Radiation == 43 {
		game.State = "LOSE"
		game.Sectors = []Sector{}
		println("Game Over!")
	}
	if sector.Radiation == 0 {
		println("Revealing adjacent sectors...")
		game.revealAdjacentSectorsTo(sector.X, sector.Y, sector.Z)
	}
}

func (game *Game) Mark(sectorId int) {
	if game.State == "LOSE" || sectorId < 0 || sectorId >= len(game.Sectors) {
		return
	}
	sector := &game.Sectors[sectorId]
	if sector.Radiation != -1 {
		return
	}
	sector.Marked = !sector.Marked
	if game.allMinesMarked() {
		game.State = "WIN"
		game.revealAllSectors()
	}
	sector.print()
}

func (game *Game) revealAllSectors() {
	for i, sector := range game.Sectors {
		game.Sectors[i].Radiation = CalculateRadiationFor(sector.X, sector.Y, sector.Z, game.mines)
	}
}

func (game Game) allMinesMarked() bool {
	markedSectors := game.getMarkedSectors()
	if len(markedSectors) != len(game.mines) {
		return false
	}
	for _, sector := range markedSectors {
		if !game.isMine(sector.Id) {
			return false
		}
	}
	return true
}

func (game Game) getMarkedSectors() []Sector {
	var markedSectors []Sector
	for _, sector := range game.Sectors {
		if sector.Marked {
			markedSectors = append(markedSectors, sector)
		}
	}
	return markedSectors
}

func (game Game) isMine(sectorId int) bool {
	sector := game.Sectors[sectorId]
	for _, mine := range game.mines {
		if sector.X == mine.X && sector.Y == mine.Y && sector.Z == mine.Z {
			return true
		}
	}
	return false
}

func (game *Game) revealAdjacentSectorsTo(x int, y int, z int) {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := z - 1; k <= z+1; k++ {
				game.reveal(i, j, k)
			}
		}
	}
}

func GetAdjacentSectorIdsFor(x int, y int, z int, scale int) []int {
	var sectorIds []int
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := z - 1; k <= z+1; k++ {
				if !(x == i && y == j && z == k) && isValidSectorLocation(i, j, k, scale) {
					sectorIds = append(sectorIds, getSectorIdFromLocation(i, j, k, scale))
				}
			}
		}
	}
	return sectorIds
}

func (game *Game) reveal(x int, y int, z int) {
	if !isValidSectorLocation(x, y, z, game.Size) {
		return
	}
	sectorId := getSectorIdFromLocation(x, y, z, game.Size)
	game.Reveal(sectorId)
}

func isValidSectorLocation(x int, y int, z int, scale int) bool {
	return isValidCoordinate(x, scale) && isValidCoordinate(y, scale) && isValidCoordinate(z, scale)
}

func isValidCoordinate(n int, scale int) bool {
	return n > -1 && n < scale
}

func getSectorIdFromLocation(x int, y int, z int, scale int) int {
	return x*int(math.Pow(float64(scale), 2)) + y*scale + z
}
