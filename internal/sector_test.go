package internal

import "testing"

func TestCalculateRadiation(t *testing.T) {
	// Given
	mines := []Location{{x: 2, y: 2, z: 2}}
	location := Location{x: 0, y: 0, z: 0}

	// When
	radiation := CalculateRadiationFor(location, mines)

	// Then
	if radiation != 0 {
		t.Errorf("Radiation should have been %d but was %d", 07, radiation)
	}
}
