package features

import (
	"net/http"
	"strconv"
	"strings"
	"sync"

	"groupieTracker/entities"
)

// Handler is the main router function for the application. It processes incoming
// HTTP requests and directs them to the appropriate handler based on the URL path.
// It handles the root path ("/") for GET requests, routes starting with "/artist/"
// to the ArtistHandler, and returns a 404 error for any other paths.
func Handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if r.Method == http.MethodGet {
		if path == "/" {
			HomeHandler(w, r)
			return
		}
		if strings.HasPrefix(path, "/artist/") {
			ArtistHandler(w, r)
			return
		}
		ErrorHandler(w, http.StatusNotFound)
	} else {
		ErrorHandler(w, http.StatusMethodNotAllowed)
	}
}

// HomeHandler processes requests for the home page. It fetches API data,
// retrieves artist information concurrently, and then renders the index.html
// template with the fetched artist data.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	allApi, err := GetApis(w, r)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return

	}

	var wg sync.WaitGroup
	artist := []entities.Artist{}
	wg.Add(1)
	go FetchData(&wg, allApi.Artists, &artist, w)
	wg.Wait()

	OpenHtml("index.html", w, artist)
}

// ArtistHandler manages requests for individual artist pages. It extracts the
// artist ID from the URL, verifies its existence, fetches detailed artist data
// and related information (locations, concert dates, and relations) concurrently,
// and then renders the about.html template with the compiled page data.
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	urlId := r.URL.Path[len("/artist/"):]
	id, err := strconv.Atoi(urlId)
	if err != nil {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if !IsArtistIDExistAPI("https://groupietrackers.herokuapp.com/api", id) {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	artist := entities.Artist{}

	apiUrl := "https://groupietrackers.herokuapp.com/api/artists/" + urlId
	var wg1 sync.WaitGroup
	wg1.Add(1)
	go FetchData(&wg1, apiUrl, &artist, w)
	wg1.Wait()

	var wg sync.WaitGroup
	pageData := &entities.PageData{Artist: artist}
	if artist.Locations != "" || artist.ConcertDates != "" || artist.Relations != "" {
		wg.Add(3)

		go FetchData(&wg, artist.Locations, &pageData.Locations, w)
		go FetchData(&wg, artist.ConcertDates, &pageData.Dates, w)
		go FetchData(&wg, artist.Relations, &pageData.Relations, w)
		
		wg.Wait()
		OpenHtml("about.html", w, pageData)

	}
}
