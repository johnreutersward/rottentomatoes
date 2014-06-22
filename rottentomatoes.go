/*
Package rottentomatoes provides a client for using the Rotten Tomatoes API.


Usage:

	import "github.com/rojters/rottentomatoes"

Construct a new Rotten Tomatoes client, then use the various services on the client to
access different parts of the Rotten Tomatoes API. For example, to find information
about a particular movie using it's IMDb Id:

	rt := rottentomatoes.NewClient(nil, "")
	// http://www.imdb.com/title/tt0118715/ (The Big Lebowski)
	m, _ := rt.DetailedInfo.MovieAlias("0118715")

Some API methods have optional parameters that can be passed. For example,
to page limit the number of results returned:

	rt := rottentomatoes.NewClient(nil, "")
	opt := &rottentomatoes.Options{PageLimit: 10}
	it, _ := rt.MovieLists.InTheaters(Opt)

The services of a client divide the API into logical chunks and correspond to
the structure of the Rotten Tomatoes API Documentation at
http://developer.rottentomatoes.com/.
*/
package rottentomatoes

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/google/go-querystring/query"
)

const (
	libraryVersion = "0.1"
	baseURL        = "http://api.rottentomatoes.com/api/public/v1.0/"
	userAgent      = "rottentomatoes/" + libraryVersion
)

// Rotten Tomatoes API client.
type Client struct {
	client *http.Client

	// ApiKey required to use the Rotten Tomatoes API.
	ApiKey string

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent used when communicating with the Rotten Tomatoes API.
	UserAgent string

	// Services are grouped as on http://developer.rottentomatoes.com/io-docs
	MovieLists    *MovieListsService
	DVDLists      *DVDListsService
	DetailedInfo  *DetailedInfoService
	Search        *SearchService
	TopLevelLists *TopLevelListsService
}

type Options struct {
	// ApiKey required to use the Rotten Tomatoes API.
	ApiKey string `url:"apikey,omitempty"`

	// Provides localized data for the selected country
	// (ISO 3166-1 alpha-2) if available. Otherwise, returns US data.
	Country string `url:"country,omitempty"`

	// Movie ID.
	Id string `url:"id,omitempty"`

	// Limit the number of entities returned
	Limit int `url:"limit,omitempty"`

	// The selected page of entities
	Page int `url:"page,omitempty"`

	// The amount of entities to get per request
	PageLimit int `url:"page_limit,omitempty"`

	//The plain text search query to search for a movie.
	Query string `url:"q,omitempty"`

	// 3 different review types are possible: 'all', 'top_critic' and 'dvd'.
	// 'top_critic' shows all the Rotten Tomatoes designated top critics.
	// 'dvd' pulls the reviews given on the DVD of the movie.
	// 'all' as the name implies retrieves all reviews.
	ReviewType string `url:"review_type,omitempty"`

	// Movies Alias type, only supports 'imdb' lookup at this time.
	Type string `url:"type,omitempty"`
}

// NewClient returns a new Rotten Tomatoes API client. If a nil
// httpClient is provided, http.DefaultClient will be used.
// A Rotten Tomatoes API key is required for use which can be
// obtained at http://developer.rottentomatoes.com/
// Provide the key as a string, or store it in an environment
// variable 'ROTTENTOMATOES_APIKEY' and pass the empty string.
func NewClient(httpClient *http.Client, apiKey string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if len(apiKey) == 0 {
		apiKey = os.Getenv("ROTTENTOMATOES_APIKEY")
	}

	baseURL, _ := url.Parse(baseURL)

	c := &Client{
		ApiKey:    apiKey,
		BaseURL:   baseURL,
		client:    httpClient,
		UserAgent: userAgent,
	}
	c.MovieLists = &MovieListsService{client: c}
	c.DVDLists = &DVDListsService{client: c}
	c.DetailedInfo = &DetailedInfoService{client: c}
	c.Search = &SearchService{client: c}
	c.TopLevelLists = &TopLevelListsService{client: c}

	return c
}

func (c *Client) get(endp string, opt *Options, r interface{}) (*http.Response, error) {

	if opt == nil {
		opt = &Options{}
	}

	opt.ApiKey = c.ApiKey

	v, err := query.Values(opt)
	if err != nil {
		return nil, err
	}

	rel, err := url.Parse(endp + v.Encode())
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotModified {
		return nil, errors.New("api error, response code: " + strconv.Itoa(resp.StatusCode))
	}

	defer resp.Body.Close()

	if r != nil {
		err = json.NewDecoder(resp.Body).Decode(r)
	}

	return resp, err

}
