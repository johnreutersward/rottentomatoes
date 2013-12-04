package rottentomatoes

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"text/template"
)

type Movie struct {
	Id               int      `json:"id"`
	Title            string   `json:"title"`
	Year             int      `json:"year"`
	Genres           []string `json:"genres"`
	MPAA_Rating      string   `json:"mpaa_rating"`
	Runtime          int      `json:"runtime"`
	CriticsConsensus string   `json:"critics_consensus"`
}

type ReleaseDates struct {
	Theater string `json:"theater"`
	DVD     string `json:"dvd"`
}

type Ratings struct {
	CriticsRating  string `json:"critics_rating"`
	CriticsScore   int    `json:"critics_score"`
	AudienceRating string `json:"audience_rating"`
	AudienceScore  int    `json:"audience_score"`
}

func UnmarshalMoviesInfo(data []byte) (Movie, error) {
	var m Movie
	//m := new(Movie)
	err := json.Unmarshal(data, &m)

	if err != nil {
		log.Panic(err)
	}

	return m, err
}

func (c *Client) MoviesInfo(id string) (Movie, error) {

	var movie Movie

	if len(c.ApiKey) == 0 {
		return movie, errors.New("missing ApiKey")
	}

	t, _ := template.New("MoviesInfoUrl").Parse(c.BaseUrl["MoviesInfo"])
	buf := new(bytes.Buffer)
	t.Execute(buf, id)

	v := url.Values{}
	v.Set("apikey", c.ApiKey)

	endp := buf.String() + v.Encode()

	data, err := c.Request(endp)

	if err != nil {
		return movie, err
	}

	movie, err = UnmarshalMoviesInfo(data)

	return movie, err
}
