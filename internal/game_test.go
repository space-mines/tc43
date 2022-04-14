package internal

import "testing"

func TestRevealSector(t *testing.T) {
	// Given
	game := NewGame("testing", []Location{{X: 2, Y: 2, Z: 2}}, nil, 3)

	// When
	game.Reveal("0")

	// Then
	size := len(game.Sectors)

	if size != 7 {
		t.Errorf("Game sector count should have been %d but was %d", 7, size)
	}
}
