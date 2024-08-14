package entities

// AllApi represents the structure of the main API response, containing URLs
// for different data endpoints.
type AllApi struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}
// Artist represents an individual artist's data, including basic information
// and URLs for additional details like locations, concert dates, and relations.
type Artist struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Image          string   `json:"image"`
	FirstAlbumDate string   `json:"firstAlbum"`
	CreationDate   int      `json:"creationDate"`
	Members        []string `json:"members"`
	Locations      string   `json:"locations"`    // api link about artist location
	ConcertDates   string   `json:"concertDates"` // api link about artist ConcertDates
	Relations      string   `json:"relations"`    // api link about artist relation
}
// Location represents location data for an artist, including an ID and a list
// of location strings.
type Location struct {
	ID        int      `json:"id"`
	Location []string `json:"locations"`
}

// Date represents concert date data for an artist, including an ID and a list
// of date strings.
type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Relation represents the relationship between dates and locations for an artist's
// performances.
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// PageData is a composite structure that combines an Artist with their associated
// Location, Date, and Relation data, used for rendering detailed artist pages.
type PageData struct {
    Artist   Artist
    Locations Location
    Dates    Date
    Relations Relation
}