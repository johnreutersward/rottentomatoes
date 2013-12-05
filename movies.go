package rottentomatoes

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/url"
	"text/template"
)

type Movie struct {
	Id                int          `json:"id"`
	Title             string       `json:"title"`
	Year              int          `json:"year"`
	Genres            []string     `json:"genres"`
	MPAA_Rating       string       `json:"mpaa_rating"`
	Runtime           int          `json:"runtime"`
	CriticsConsensus  string       `json:"critics_consensus"`
	ReleaseDates      ReleaseDates `json:"release_dates"`
	Ratings           Ratings      `json:"ratings"`
	Synopsis          string       `json:"synopsis"`
	Posters           Posters      `json:"posters"`
	AbridgedCast      []Cast       `json:"abridged_cast"`
	AbridgedDirectors []Director   `json:"abridged_directors"`
	Studio            string       `json:"studio"`
	AlternateIds      AlternateIds `json:"alternate_ids"`
	Links             Links        `json:"links"`
}

type Movies struct {
	Id                string       `json:"id"`
	Title             string       `json:"title"`
	Year              int          `json:"year"`
	Genres            []string     `json:"genres"`
	MPAA_Rating       string       `json:"mpaa_rating"`
	Runtime           int          `json:"runtime"`
	CriticsConsensus  string       `json:"critics_consensus"`
	ReleaseDates      ReleaseDates `json:"release_dates"`
	Ratings           Ratings      `json:"ratings"`
	Synopsis          string       `json:"synopsis"`
	Posters           Posters      `json:"posters"`
	AbridgedCast      []Cast       `json:"abridged_cast"`
	AbridgedDirectors []Director   `json:"abridged_directors"`
	Studio            string       `json:"studio"`
	AlternateIds      AlternateIds `json:"alternate_ids"`
	Links             Links        `json:"links"`
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

type Posters struct {
	Thumbnail string `json:"thumbnail"`
	Profile   string `json:"Profile"`
	Detailed  string `json:"Detailed"`
	Original  string `json:"original"`
}

type Director struct {
	Name string `json:"name"`
}

type AlternateIds struct {
	IMDB string `json:"imdb"`
}

type Links struct {
	Self      string `json:"self"`
	Alternate string `json:"alternate"`
	Cast      string `json:"cast"`
	Clips     string `json:"clips"`
	Reviews   string `json:"reviews"`
	Similar   string `json:"similar"`
}

func UnmarshalMoviesInfo(data []byte) (Movie, error) {
	var m Movie
	err := json.Unmarshal(data, &m)

	return m, err
}

func UnmarshalMovies(data []byte) ([]Movies, error) {
	var jsond map[string]*json.RawMessage
	_ = json.Unmarshal(data, &jsond)

	var movieList []Movies
	err := json.Unmarshal(*jsond["movies"], &movieList)

	return movieList, err
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

func (c *Client) MoviesSimilar(id string, limit string) ([]Movies, error) {

	var movies []Movies

	if len(c.ApiKey) == 0 {
		return movies, errors.New("missing ApiKey")
	}

	t, _ := template.New("MoviesSimilarUrl").Parse(c.BaseUrl["MoviesSimilar"])
	buf := new(bytes.Buffer)
	t.Execute(buf, id)

	v := url.Values{}
	v.Set("limit", limit)
	v.Set("apikey", c.ApiKey)

	endp := buf.String() + v.Encode()

	data, err := c.Request(endp)

	if err != nil {
		return movies, err
	}

	movies, err = UnmarshalMovies(data)

	return movies, err
}

func (c *Client) MoviesAlias(id string) (Movie, error) {

	var movie Movie

	if len(c.ApiKey) == 0 {
		return movie, errors.New("missing ApiKey")
	}

	t, _ := template.New("MoviesAliasUrl").Parse(c.BaseUrl["MoviesAlias"])
	buf := new(bytes.Buffer)
	t.Execute(buf, id)

	v := url.Values{}
	v.Set("id", id)
	v.Set("type", "imdb")
	v.Set("apikey", c.ApiKey)

	endp := buf.String() + v.Encode()

	data, err := c.Request(endp)

	if err != nil {
		return movie, err
	}

	movie, err = UnmarshalMoviesInfo(data)

	return movie, err
}
