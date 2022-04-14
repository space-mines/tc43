package internal

type Sector struct {
	Id        int  `json:"id"`
	X         int  `json:"X"`
	Y         int  `json:"Y"`
	Z         int  `json:"Z"`
	Radiation int  `json:"radiation"`
	Marked    bool `json:"marked"`
}
