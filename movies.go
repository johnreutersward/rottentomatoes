package rottentomatoes

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
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

type Review struct {
	Critic      string      `json:"critic"`
	Date        string      `json:"date"`
	Freshness   string      `json:"freshness"`
	Publication string      `json:"publication"`
	Quote       string      `json:"quote"`
	Links       ReviewLinks `json:"links"`
}

type ReviewLinks struct {
	Review string `json:"review"`
}

type Clip struct {
	Title     string    `json:"title"`
	Duration  string    `json:"duration"`
	Thumbnail string    `json:"thumbnail"`
	Links     ClipLinks `json:"links"`
}

type ClipLinks struct {
	Alternate string `json:"alternate"`
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

func UnmarshalSearch(data []byte) ([]Movies, int, error) {
	var jsond map[string]*json.RawMessage
	_ = json.Unmarshal(data, &jsond)

	var movieList []Movies
	err := json.Unmarshal(*jsond["movies"], &movieList)

	var t int
	err = json.Unmarshal(*jsond["total"], &t)

	return movieList, t, err
}

func UnmarshalReviews(data []byte) ([]Review, int, error) {
	var jsond map[string]*json.RawMessage
	_ = json.Unmarshal(data, &jsond)

	var reviewList []Review
	err := json.Unmarshal(*jsond["reviews"], &reviewList)

	var t int
	err = json.Unmarshal(*jsond["total"], &t)

	return reviewList, t, err
}

func UnmarshalClips(data []byte) ([]Clip, error) {
	var jsond map[string]*json.RawMessage
	_ = json.Unmarshal(data, &jsond)

	var clipList []Clip
	err := json.Unmarshal(*jsond["clips"], &clipList)

	return clipList, err
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

func (c *Client) MoviesSimilar(id string, limit int) ([]Movies, error) {

	var movies []Movies

	if len(c.ApiKey) == 0 {
		return movies, errors.New("missing ApiKey")
	}

	t, _ := template.New("MoviesSimilarUrl").Parse(c.BaseUrl["MoviesSimilar"])
	buf := new(bytes.Buffer)
	t.Execute(buf, id)

	limit_t := strconv.Itoa(limit)

	v := url.Values{}
	v.Set("limit", limit_t)
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

func (c *Client) MoviesSearch(q string, page_limit int, page int) ([]Movies, int, error) {

	var movies []Movies
	var total int

	if len(c.ApiKey) == 0 {
		return movies, 0, errors.New("missing ApiKey")
	}

	t, _ := template.New("MoviesSearchUrl").Parse(c.BaseUrl["Search"])
	buf := new(bytes.Buffer)
	t.Execute(buf, q)

	page_limit_t := strconv.Itoa(page_limit)
	page_t := strconv.Itoa(page)

	v := url.Values{}
	v.Set("q", q)
	v.Set("page_limit", page_limit_t)
	v.Set("page", page_t)
	v.Set("apikey", c.ApiKey)

	endp := buf.String() + v.Encode()

	data, err := c.Request(endp)

	if err != nil {
		return movies, 0, err
	}

	movies, total, err = UnmarshalSearch(data)

	return movies, total, err
}

func (c *Client) MoviesReviews(id string, review_type string, page_limit int, page int, country string) ([]Review, int, error) {

	var reviews []Review
	var total int

	if len(c.ApiKey) == 0 {
		return reviews, 0, errors.New("missing ApiKey")
	}

	t, _ := template.New("MovieReviewsUrl").Parse(c.BaseUrl["MovieReviews"])
	buf := new(bytes.Buffer)
	t.Execute(buf, id)

	page_limit_t := strconv.Itoa(page_limit)
	page_t := strconv.Itoa(page)

	v := url.Values{}
	v.Set("review_type", review_type)
	v.Set("page_limit", page_limit_t)
	v.Set("page", page_t)
	v.Set("country", country)
	v.Set("apikey", c.ApiKey)

	endp := buf.String() + v.Encode()

	data, err := c.Request(endp)

	if err != nil {
		return reviews, 0, err
	}

	reviews, total, err = UnmarshalReviews(data)

	return reviews, total, err
}

func (c *Client) MoviesClips(id string) ([]Clip, error) {

	var clips []Clip

	if len(c.ApiKey) == 0 {
		return clips, errors.New("missing ApiKey")
	}

	t, _ := template.New("MoviesClipsUrl").Parse(c.BaseUrl["MovieClips"])
	buf := new(bytes.Buffer)
	t.Execute(buf, id)

	v := url.Values{}
	v.Set("apikey", c.ApiKey)

	endp := buf.String() + v.Encode()

	data, err := c.Request(endp)

	if err != nil {
		return clips, err
	}

	clips, err = UnmarshalClips(data)

	return clips, err
}
