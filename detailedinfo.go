package rottentomatoes

import (
	"fmt"
	"net/url"
)

type CastInfoResponse struct {
	Cast  []Cast        `json:"cast,omitempty"`
	Links CastInfoLinks `json:"links,omitempty"`
}

type Cast struct {
	Name       string   `json:"name"`
	Id         Id       `json:"id"`
	Characters []string `json:"characters"`
}

type CastInfoLinks struct {
	Rel string `json:"rel,omitempty"`
}

type MovieClipsResponse struct {
	Clips []Clip         `json:"clips,omitempty"`
	Links MovieClipLinks `json:"links,omitempty"`
}

type MovieClipLinks struct {
	Alternate string `json:"alternate,omitempty"`
	Rel       string `json:"rel,omitempty"`
	Self      string `json:"self,omitempty"`
}

type Clip struct {
	Title     string    `json:"title,omitempty"`
	Duration  string    `json:"duration,omitempty"`
	Thumbnail string    `json:"thumbnail,omitempty"`
	Links     ClipLinks `json:"links,omitempty"`
}

type ClipLinks struct {
	Alternate string `json:"alternate,omitempty"`
}

type MovieReviewsResponse struct {
	Total        int               `json:"total,omitempty"`
	Reviews      []Review          `json:"reviews,omitempty"`
	Links        MovieReviewsLinks `json:"links,omitempty"`
	LinkTemplate string            `json:"link_template,omitempty"`
}

type MovieReviewsLinks struct {
	Alternate string `json:"alternate,omitempty"`
	Next      string `json:"next,omitempty"`
	Rel       string `json:"rel,omitempty"`
	Self      string `json:"self,omitempty"`
}

type Review struct {
	Critic      string      `json:"critic,omitempty"`
	Date        string      `json:"date,omitempty"`
	Freshness   string      `json:"freshness,omitempty"`
	Publication string      `json:"publication,omitempty"`
	Quote       string      `json:"quote,omitempty"`
	Links       ReviewLinks `json:"links,omitempty"`
}

type ReviewLinks struct {
	Review string `json:"review,omitempty"`
}

type MovieSimilarResponse struct {
	Movies       []Movie           `json:"movies,omitempty"`
	Links        MovieSimilarLinks `json:"links,omitempty"`
	LinkTemplate string            `json:"link_template,omitempty"`
}

type MovieSimilarLinks struct {
	Rel  string `json:"rel,omitempty"`
	Self string `json:"self,omitempty"`
}

type DetailedInfoService struct {
	client *Client
}

// Returns detailed information on a specific movie specified by Id.
func (d *DetailedInfoService) MovieInfo(id string) (*Movie, error) {
	rel := fmt.Sprintf("movies/%s.json?", url.QueryEscape(id))
	resp := new(Movie)
	_, err := d.client.get(rel, nil, resp)
	return resp, err
}

// Returns the complete movie cast for a movie.
func (d *DetailedInfoService) CastInfo(id string) (*CastInfoResponse, error) {
	rel := fmt.Sprintf("movies/%s/cast.json?", url.QueryEscape(id))
	resp := new(CastInfoResponse)
	_, err := d.client.get(rel, nil, resp)
	return resp, err
}

// Returns related movie clips and trailers for a movie
func (d *DetailedInfoService) MovieClips(id string) (*MovieClipsResponse, error) {
	rel := fmt.Sprintf("movies/%s/clips.json?", url.QueryEscape(id))
	resp := new(MovieClipsResponse)
	_, err := d.client.get(rel, nil, resp)
	return resp, err
}

// Returns the reviews for a movie. Results are paginated if they go past the specified page limit.
// Available options: PageLimit, Page, Country, ReviewType
func (d *DetailedInfoService) MovieReviews(id string, opt *Options) (*MovieReviewsResponse, error) {
	rel := fmt.Sprintf("movies/%s/reviews.json?", url.QueryEscape(id))
	resp := new(MovieReviewsResponse)
	_, err := d.client.get(rel, opt, resp)
	return resp, err
}

// Returns similar movies for a movie.
// Available options: Limit
func (d *DetailedInfoService) MovieSimilar(id string, opt *Options) (*MovieSimilarResponse, error) {
	rel := fmt.Sprintf("movies/%s/similar.json?", url.QueryEscape(id))
	resp := new(MovieSimilarResponse)
	_, err := d.client.get(rel, opt, resp)
	return resp, err
}

// Same as MovieInfo but provides a movie lookup by an id from a different vendor.
// Only supports imdb lookup at this time. Note: skip the initial 'tt'.
func (d *DetailedInfoService) MovieAlias(imdbId string) (*Movie, error) {
	rel := "movie_alias.json?"
	opt := &Options{
		Id:   url.QueryEscape(imdbId),
		Type: "imdb",
	}
	resp := new(Movie)
	_, err := d.client.get(rel, opt, resp)
	return resp, err
}
