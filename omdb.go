package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// Rating describes the Rating object of a movie
type Rating struct {
	Source string
	Value  string
}

// Movie describes a movie response
type Movie struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Ratings    []Rating
	Metascore  string
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
	Response   string
}

// buildURLwithParams builds the URL parameters
func buildURLwithParams(name string, url string) string {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalf("fetch: %v", err)
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("OMDB_KEY"))
	q.Add("t", name)
	q.Add("plot", "full")
	req.URL.RawQuery = q.Encode()

	return req.URL.String()
}

// getMovieResponse gets the JSON of the requested movie
func getMovieResponse(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("fetch: %v", err)
	}
	log.Println(url)
	log.Println(resp)
	movie, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalf("fetch: reading %s: %v\n", url, err)
	}
	return movie
}

// unmarshalMovieIntoStruct puts the JSON response into a Movie struct
func unmarshalMovieIntoStruct(movie []byte) Movie {
	var m Movie
	if err := json.Unmarshal(movie, &m); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	return m
}

// handlePoster puts the movie poster into the writer response
func handlePoster(w http.ResponseWriter, r *http.Request) {
	title := strings.TrimPrefix(r.URL.Path, "/movie/")
	movieResponse := getMovieResponse(buildURLwithParams(title, os.Getenv("OMDB_URL")))
	movie := unmarshalMovieIntoStruct(movieResponse)
	fmt.Fprintf(w, fmt.Sprintf("<h1>%s (%s)</h1>", movie.Title, movie.Year))
	fmt.Fprintf(w, fmt.Sprintf("<img src=\"%s\"/>", movie.Poster))
}

func main() {
	http.HandleFunc("/movie/", handlePoster)
	http.ListenAndServe("0.0.0.0:8080", nil)
	log.Println("Server running...")
}
