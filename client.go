// Package rottentomatoes implements a wrapper for the Rotten Tomatoes Web Api.
package rottentomatoes

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

type client struct {
	ApiKey  string
	BaseUrl map[string]string
}

type apiError struct {
	Error string `json:"error"`
}

// NewClient creates a rottentomatoes client instance.
func NewClient() (c *client, err error) {

	apikey := os.Getenv("ROTTENTOMATOES_APIKEY")

	if len(apikey) == 0 {
		err = errors.New("missing api key")
		return
	}

	c = &client{
		ApiKey: apikey,
		BaseUrl: map[string]string{
			"MoviesInfo":    "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}.json?",
			"CastInfo":      "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/cast.json?",
			"MovieClips":    "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/clips.json?",
			"MovieReviews":  "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/reviews.json?",
			"MoviesSimilar": "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/similar.json?",
			"MoviesAlias":   "http://api.rottentomatoes.com/api/public/v1.0/movie_alias.json?",
			"Search":        "http://api.rottentomatoes.com/api/public/v1.0/movies.json?",
		},
	}
	return
}

func (c *client) request(endp string) (data []byte, err error) {

	resp, err := http.Get(endp)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, _ = ioutil.ReadAll(resp.Body)

	var e apiError
	_ = json.Unmarshal(data, &e)

	if len(e.Error) != 0 {
		return nil, errors.New(e.Error)
	}

	return data, nil

}
