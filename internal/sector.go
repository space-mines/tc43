package internal

type Sector struct {
	Id        int  `json:"id"`
	X         int  `json:"x"`
	Y         int  `json:"y"`
	Z         int  `json:"z"`
	Radiation int  `json:"radiation"`
	Marked    bool `json:"marked"`
}

func RemoveSectorsWithNoRadiation(sectors []Sector) []Sector {
	clean := make([]Sector, 0)
	for _, sector := range sectors {
		if sector.Radiation != 0 {
			clean = append(clean, sector)
		}
	}
	return clean
}
