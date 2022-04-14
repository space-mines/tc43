package internal

type Sector struct {
	Id        int  `json:"id"`
	X         int  `json:"x"`
	Y         int  `json:"y"`
	Z         int  `json:"z"`
	Radiation int  `json:"radiation"`
	Marked    bool `json:"marked"`
}
