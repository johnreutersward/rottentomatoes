// Package rottentomatoes implements a wrapper for the Rotten Tomatoes Web Api.
package rottentomatoes

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

type Client struct {
	ApiKey  string
	BaseUrl map[string]string
}

type apiError struct {
	Error string `json:"error"`
}

// NewClient creates a rottentomatoes client instance.
func NewClient() (c *Client, err error) {

	apikey := os.Getenv("ROTTENTOMATOES_APIKEY")

	if len(apikey) == 0 {
		err = errors.New("missing api key")
		return
	}

	c = &Client{
		ApiKey: apikey,
		BaseUrl: map[string]string{
			"BoxOfficeMovies":  "http://api.rottentomatoes.com/api/public/v1.0/lists/movies/box_office.json?",
			"InTheatersMovies": "http://api.rottentomatoes.com/api/public/v1.0/lists/movies/in_theaters.json?",
			"MovieInfo":        "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}.json?",
			"MovieCast":        "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/cast.json?",
			"MovieClips":       "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/clips.json?",
			"MovieReviews":     "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/reviews.json?",
			"MovieSimilar":     "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/similar.json?",
			"MovieAlias":       "http://api.rottentomatoes.com/api/public/v1.0/movie_alias.json?",
			"MoviesSearch":     "http://api.rottentomatoes.com/api/public/v1.0/movies.json?",
		},
	}
	return
}

func (c *Client) request(endp string) (data []byte, err error) {

	resp, err := http.Get(endp)

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
