package groupie

var (
	Artist      Artists
	Location    Locations
	ConcertDate ConcertDates
	Relation    Relations
	Cards       Card
	Fetched     bool
)

type Artists []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}
type Locations struct {
	Locations []string `json:"locations"`
}
type ConcertDates struct {
	Dates []string `json:"dates"`
}

type Relations struct {
	Relation map[string][]string `json:"datesLocations"`
}
type Card struct {
	Art  Artists
	Loca Locations
	Conc ConcertDates
	Rela Relations
}

const Url = "https://groupietrackers.herokuapp.com/api/"
