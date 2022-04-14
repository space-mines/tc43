package internal

import "testing"

func TestRevealSector(t *testing.T) {
	// Given
	game := NewGame("testing", []Location{{x: 2, y: 2, z: 2}}, nil, 3)

	// When
	game.Reveal("0")

	// Then
	size := len(game.Sectors)

	if size != 7 {
		t.Errorf("Game sector count should have been %d but was %d", 7, size)
	}
}
