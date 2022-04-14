package internal

type Location struct {
	X int
	Y int
	Z int
}

func contains(locations []Location, x int, y int, z int) bool {
	for _, location := range locations {
		if location.X == x && location.Y == y && location.Z == z {
			return true
		}
	}
	return false
}

func countAdjacentMines(locations []Location, x int, y int, z int) int {
	count := 0
	for _, location := range locations {
		if location.X <= x+1 &&
			location.X >= x-1 &&
			location.Y <= y+1 &&
			location.Y >= y-1 &&
			location.Z <= z+1 &&
			location.Z >= z-1 {
			count++
		}
	}
	return count
}

func (location Location) CalculateRadiation(mines []Location) int {
	if contains(mines, location.X, location.Y, location.Z) {
		return 43
	}
	return countAdjacentMines(mines, location.X, location.Y, location.Z)
}
