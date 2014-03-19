// Package rottentomatoes implements a wrapper for the Rotten Tomatoes Web Api.
//
// Example:
//
//	package main
//
//	import (
//		"fmt"
//		"log"
//		"net/http"
//
//		"github.com/rojters/rottentomatoes"
//	)
//
//	func main() {
//		rt, _ := rottentomatoes.NewClient(&http.Client{})
//
//		// Get info using movie id
//		m, err := rt.MovieInfo("14281")
//
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		fmt.Printf("Title: %s, Year: %d, Runtime: %d min\n", m.Title, m.Year, m.Runtime)
//		// Title: The Big Lebowski, Year: 1998, Runtime: 118 min
//
//	}
//
package rottentomatoes

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

type Client struct {
	client  *http.Client
	ApiKey  string
	BaseUrl map[string]string
}

type apiError struct {
	Error string `json:"error"`
}

// NewClient creates a rottentomatoes client instance.
func NewClient(httpClient *http.Client) (c *Client, err error) {

	apikey := os.Getenv("ROTTENTOMATOES_APIKEY")

	if len(apikey) == 0 {
		err = errors.New("missing api key")
		return
	}

	c = &Client{
		client: httpClient,
		ApiKey: apikey,
		BaseUrl: map[string]string{
			"BoxOfficeMovies":    "http://api.rottentomatoes.com/api/public/v1.0/lists/movies/box_office.json?",
			"InTheatersMovies":   "http://api.rottentomatoes.com/api/public/v1.0/lists/movies/in_theaters.json?",
			"OpeningMovies":      "http://api.rottentomatoes.com/api/public/v1.0/lists/movies/opening.json?",
			"UpcomingMovies":     "http://api.rottentomatoes.com/api/public/v1.0/lists/movies/upcoming.json?",
			"TopRentals":         "http://api.rottentomatoes.com/api/public/v1.0/lists/dvds/top_rentals.json?",
			"CurrentReleaseDVDs": "http://api.rottentomatoes.com/api/public/v1.0/lists/dvds/current_releases.json?",
			"NewReleaseDVDs":     "http://api.rottentomatoes.com/api/public/v1.0/lists/dvds/new_releases.json?",
			"UpcomingDVDs":       "http://api.rottentomatoes.com/api/public/v1.0/lists/dvds/upcoming.json?",
			"MovieInfo":          "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}.json?",
			"MovieCast":          "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/cast.json?",
			"MovieClips":         "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/clips.json?",
			"MovieReviews":       "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/reviews.json?",
			"MovieSimilar":       "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/similar.json?",
			"MovieAlias":         "http://api.rottentomatoes.com/api/public/v1.0/movie_alias.json?",
			"MoviesSearch":       "http://api.rottentomatoes.com/api/public/v1.0/movies.json?",
		},
	}
	return
}

func (c *Client) request(endp string) (data []byte, err error) {

	resp, err := c.client.Get(endp)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	data, _ = ioutil.ReadAll(resp.Body)

	var e apiError
	_ = json.Unmarshal(data, &e)

	if len(e.Error) != 0 {
		err = errors.New(e.Error)
		return
	}

	return
}
