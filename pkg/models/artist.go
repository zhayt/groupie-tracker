package models

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type PresentAllData struct {
	Artist       *Artist
	Locations    *Locations
	ConcertDates *ConcertDates
	Relations    *Relations
}

type Locations struct {
	Locations []string `json:"locations"`
}

type ConcertDates struct {
	Dates []string `json:"dates"`
}

type Relations struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

func New() *PresentAllData {
	return &PresentAllData{
		Artist:       &Artist{},
		Locations:    &Locations{},
		ConcertDates: &ConcertDates{},
		Relations:    &Relations{},
	}
}
