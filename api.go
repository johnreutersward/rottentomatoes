package rottentomatoes

import (
	"bytes"
	"net/url"
	"strconv"
	"text/template"
)

func (c *client) MovieInfo(id string) (movie Movie, err error) {

	endpoint := c.getEndpoint("MovieInfo", id)
	urlParams := c.prepareUrl(nil)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	movie, err = UnmarshalMoviesInfo(data)
	return
}

func (c *client) MovieCast(id string) (castList []Cast, err error) {

	endpoint := c.getEndpoint("MovieCast", id)
	urlParams := c.prepareUrl(nil)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	castList, err = UnmarshalCastInfo(data)
	return
}

func (c *client) MovieClips(id string) (clips []Clip, err error) {

	endpoint := c.getEndpoint("MovieClips", id)
	urlParams := c.prepareUrl(nil)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	clips, err = UnmarshalClips(data)
	return
}

func (c *client) MoviesReviews(id string, review_type string, page_limit int, page int, country string) ([]Review, int, error) {

	var reviews []Review
	var total int

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

	data, err := c.request(endp)

	if err != nil {
		return reviews, 0, err
	}

	reviews, total, err = UnmarshalReviews(data)

	return reviews, total, err
}

func (c *client) MoviesSimilar(id string, limit int) ([]Movie_, error) {

	var movies []Movie_

	t, _ := template.New("MoviesSimilarUrl").Parse(c.BaseUrl["MoviesSimilar"])
	buf := new(bytes.Buffer)
	t.Execute(buf, id)

	limit_t := strconv.Itoa(limit)

	v := url.Values{}
	v.Set("limit", limit_t)
	v.Set("apikey", c.ApiKey)

	endp := buf.String() + v.Encode()

	data, err := c.request(endp)

	if err != nil {
		return movies, err
	}

	movies, err = UnmarshalMovies(data)

	return movies, err
}

func (c *client) MoviesAlias(id string) (Movie, error) {

	var movie Movie

	t, _ := template.New("MoviesAliasUrl").Parse(c.BaseUrl["MoviesAlias"])
	buf := new(bytes.Buffer)
	t.Execute(buf, id)

	v := url.Values{}
	v.Set("id", id)
	v.Set("type", "imdb")
	v.Set("apikey", c.ApiKey)

	endp := buf.String() + v.Encode()

	data, err := c.request(endp)

	if err != nil {
		return movie, err
	}

	movie, err = UnmarshalMoviesInfo(data)

	return movie, err
}

func (c *client) MoviesSearch(q string, page_limit int, page int) ([]Movie_, int, error) {

	var movies []Movie_
	var total int

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

	data, err := c.request(endp)

	if err != nil {
		return movies, 0, err
	}

	movies, total, err = UnmarshalSearch(data)

	return movies, total, err
}
