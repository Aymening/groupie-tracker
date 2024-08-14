package features

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"groupieTracker/entities"
)
// GetApis fetches the main API data from the Groupie Trackers API endpoint.
// It returns an AllApi struct containing information about various API endpoints,
// or an error if the request fails or the response cannot be decoded.
func GetApis(w http.ResponseWriter, r *http.Request) (entities.AllApi, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		fmt.Printf("error fetching main API: %v", err)
		return entities.AllApi{}, err
	}
	defer resp.Body.Close()

	var allApi entities.AllApi
	if err := json.NewDecoder(resp.Body).Decode(&allApi); err != nil {
		fmt.Printf("error decoding main API: %v", err)
		return entities.AllApi{}, err
	}

	return allApi, nil
}
// FetchData is a concurrent function that retrieves data from a given URL and
// decodes it into the provided target interface. It's designed to be used with
// a WaitGroup for synchronization. If an error occurs during the request or
// decoding, it calls the ErrorHandler function.
func FetchData(wg *sync.WaitGroup, url string, target interface{}, w http.ResponseWriter) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(target)
}
// IsArtistIDExistAPI checks if an artist with the given ID exists in the API.
// It constructs a URL using the provided base URL and artist ID, makes a GET
// request, and attempts to decode the response into an Artist struct.
// Returns true if the artist exists (i.e., has a non-zero ID), false otherwise
// or if any errors occur during the process.
func IsArtistIDExistAPI(apiBaseURL string, id int) bool {
	url := fmt.Sprintf("%s/artists/%d", apiBaseURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	var artist entities.Artist
	if err := json.NewDecoder(resp.Body).Decode(&artist); err != nil {
		return false
	}
	if artist.ID == 0 {
		return false
	}

	return true 
}
