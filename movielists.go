package rottentomatoes

import (
	"encoding/json"
	"strconv"
	"strings"
)

type MovieListsResponse struct {
	Total        int        `json:"total,omitempty"`
	Movies       []Movie    `json:"movies,omitempty"`
	Links        MovieLinks `json:"links,omitempty"`
	LinkTemplate string     `json:"link_template,omitempty"`
}

type Movie struct {
	Id                Id           `json:"id,omitempty"`
	Title             string       `json:"title,omitempty"`
	Year              IntWrapper   `json:"year,omitempty"`
	Genres            []string     `json:"genres,omitempty"`
	MPAA_Rating       string       `json:"mpaa_rating,omitempty"`
	Runtime           IntWrapper   `json:"runtime,omitempty"`
	CriticsConsensus  string       `json:"critics_consensus,omitempty"`
	ReleaseDates      ReleaseDates `json:"release_dates,omitempty"`
	Ratings           Ratings      `json:"ratings,omitempty"`
	Synopsis          string       `json:"synopsis,omitempty"`
	Posters           Posters      `json:"posters,omitempty"`
	AbridgedCast      []Cast       `json:"abridged_cast,omitempty"`
	AbridgedDirectors []Director   `json:"abridged_directors,omitempty"`
	Studio            string       `json:"studio,omitempty"`
	AlternateIds      AlternateIds `json:"alternate_ids,omitempty"`
	Links             MovieLinks   `json:"links,omitempty"`
	LinkTemplate      string       `json:"link_template,omitempty"`
}

// The Rotten Tomatoes API sends an empty string if the field is missing rather than omitting it.
type IntWrapper int64

func (i *IntWrapper) UnmarshalJSON(b []byte) (err error) {
	var s string
	err = json.Unmarshal(b, &s)
	if err == nil {
		*i = IntWrapper(0)
		return
	}
	var n int
	err = json.Unmarshal(b, &n)
	if err == nil {
		*i = IntWrapper(n)
	}
	return
}

type Id int64

func (i *Id) UnmarshalJSON(b []byte) (err error) {
	if n, err := strconv.Atoi(strings.Trim(string(b), `"`)); err == nil {
		*i = Id(n)
	}
	return err
}

type ReleaseDates struct {
	Theater string `json:"theater,omitempty"`
	DVD     string `json:"dvd,omitempty"`
}

type Ratings struct {
	CriticsRating  string `json:"critics_rating,omitempty"`
	CriticsScore   int    `json:"critics_score,omitempty"`
	AudienceRating string `json:"audience_rating,omitempty"`
	AudienceScore  int    `json:"audience_score,omitempty"`
}

type Posters struct {
	Thumbnail string `json:"thumbnail,omitempty"`
	Profile   string `json:"Profile,omitempty"`
	Detailed  string `json:"Detailed,omitempty"`
	Original  string `json:"original,omitempty"`
}

type Director struct {
	Name string `json:"name,omitempty"`
}

type AlternateIds struct {
	IMDB Id `json:"imdb,omitempty"`
}

type MovieLinks struct {
	Self      string `json:"self,omitempty"`
	Alternate string `json:"alternate,omitempty"`
	Cast      string `json:"cast,omitempty"`
	Clips     string `json:"clips,omitempty"`
	Reviews   string `json:"reviews,omitempty"`
	Similar   string `json:"similar,omitempty"`
	Canonical string `json:"canonical,omitempty"`
}

type MovieListsService struct {
	client *Client
}

// Returns Top Box Office Earning Movies, Sorted by Most Recent Weekend Gross Ticket Sales.
// Available options: Limit, Country.
func (m *MovieListsService) BoxOffice(opt *Options) (*MovieListsResponse, error) {
	rel := "lists/movies/box_office.json?"
	resp := new(MovieListsResponse)
	_, err := m.client.get(rel, opt, resp)
	return resp, err
}

// Returns movies currently in theaters.
// Available options: PageLimit, Page, Country.
func (m *MovieListsService) InTheaters(opt *Options) (*MovieListsResponse, error) {
	rel := "lists/movies/in_theaters.json?"
	resp := new(MovieListsResponse)
	_, err := m.client.get(rel, opt, resp)
	return resp, err
}

// Returns current opening movies.
// Available options: Limit, Country.
func (m *MovieListsService) OpeningMovies(opt *Options) (*MovieListsResponse, error) {
	rel := "lists/movies/opening.json?"
	resp := new(MovieListsResponse)
	_, err := m.client.get(rel, opt, resp)
	return resp, err
}

// Returns upcoming movies. Results are paginated if they go past the specified page limit.
// Available options: PageLimit, Page, Country.
func (m *MovieListsService) UpcomingMovies(opt *Options) (*MovieListsResponse, error) {
	rel := "lists/movies/upcoming.json?"
	resp := new(MovieListsResponse)
	_, err := m.client.get(rel, opt, resp)
	return resp, err
}
