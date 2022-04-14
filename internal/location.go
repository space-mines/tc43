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

func (Location Location) CalculateRadiation(mines []Location) int {
	radiation := 0
	if contains(mines, Location.X-1, Location.Y-1, Location.Z-1) {
		radiation++
	}
	return radiation
}
