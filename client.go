package rottentomatoes

import (
	"io/ioutil"
    "net/http"
    "encoding/json"
    "errors"
    //"fmt"
)

type Client struct {
	ApiKey string
	BaseUrl map[string]string
}

type ApiError struct {
    Error string `json:"error"`
}

func NewClient(apikey string) *Client {
	c := &Client {
		ApiKey: apikey,
		BaseUrl:  map[string]string {
			"MoviesInfo": "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}.json?",
			"CastInfo": "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/cast.json?",
			"MovieClips": "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/clips.json?",
			"MovieReviews": "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/reviews.json?",
			"MoviesSimilar": "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/similar.json?",
			"MoviesAlias": "http://api.rottentomatoes.com/api/public/v1.0/movie_alias.json?",
			"Search": "http://api.rottentomatoes.com/api/public/v1.0/movies.json?",
		},
	}
	return c
}

func (c *Client) Request(endp string) ([]byte, error) {
	
	resp, err := http.Get(endp)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    data, _ := ioutil.ReadAll(resp.Body)
    
    var e ApiError
    _ = json.Unmarshal(data, &e)

    if len(e.Error) != 0 {
        return nil, errors.New(e.Error)
    }
    
    return data, nil

}