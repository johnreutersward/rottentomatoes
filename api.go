package rottentomatoes

import "strconv"

// MovieInfo, Detailed information on a specific movie specified by Id.
func (c *Client) MovieInfo(id string) (movie Movie, err error) {

	endpoint := c.getEndpoint("MovieInfo", id)
	urlParams := c.prepareUrl(nil)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	movie, err = unmarshalMoviesInfo(data)
	return
}

// MovieCast, Pulls the complete movie cast for a movie.
func (c *Client) MovieCast(id string) (castList []Cast, err error) {

	endpoint := c.getEndpoint("MovieCast", id)
	urlParams := c.prepareUrl(nil)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	castList, err = unmarshalCastInfo(data)
	return
}

// MovieClips, Related movie clips and trailers for a movie.
func (c *Client) MovieClips(id string) (clips []Clip, err error) {

	endpoint := c.getEndpoint("MovieClips", id)
	urlParams := c.prepareUrl(nil)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	clips, err = unmarshalClips(data)
	return
}

// MovieReviews, Retrieves the reviews for a movie. Results are paginated if they go past the specified page limit.
func (c *Client) MovieReviews(id string, review_type string, page_limit int, page int, country string) (reviews []Review, total int, err error) {

	endpoint := c.getEndpoint("MovieReviews", id)
	page_limit_t := strconv.Itoa(page_limit)
	page_t := strconv.Itoa(page)

	q := map[string]string{
		"review_type": review_type,
		"page_limit":  page_limit_t,
		"page":        page_t,
		"country":     country,
	}

	urlParams := c.prepareUrl(q)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	reviews, total, err = unmarshalReviews(data)
	return
}

// MovieSimilar returns similar movies.
func (c *Client) MovieSimilar(id string, limit int) (movies []Movie_, err error) {

	endpoint := c.getEndpoint("MovieSimilar", id)
	limit_t := strconv.Itoa(limit)

	q := map[string]string{
		"limit": limit_t,
	}

	urlParams := c.prepareUrl(q)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	movies, err = unmarshalMovies(data)
	return
}

// MovieAlias, Provides a movie lookup by an id from a different vendor. Only supports imdb lookup at this time.
func (c *Client) MovieAlias(id string) (movie Movie, err error) {

	endpoint := c.getEndpoint("MovieAlias", id)

	q := map[string]string{
		"id":   id,
		"type": "imdb",
	}

	urlParams := c.prepareUrl(q)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	movie, err = unmarshalMoviesInfo(data)
	return
}

// MoviesSearch returns a list movies that matches the query string.
func (c *Client) MoviesSearch(q string, page_limit int, page int) (movies []Movie_, total int, err error) {

	endpoint := c.getEndpoint("MoviesSearch", "")

	page_limit_t := strconv.Itoa(page_limit)
	page_t := strconv.Itoa(page)

	queryParams := map[string]string{
		"q":          q,
		"page_limit": page_limit_t,
		"page":       page_t,
	}

	urlParams := c.prepareUrl(queryParams)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	movies, total, err = unmarshalSearch(data)
	return
}

// BoxOfficeMovies, Top Box Office Earning Movies, Sorted by Most Recent Weekend Gross Ticket Sales.
func (c *Client) BoxOfficeMovies(limit int, country string) (movies []Movie_, err error) {

	endpoint := c.getEndpoint("BoxOfficeMovies", "")
	limit_t := strconv.Itoa(limit)

	q := map[string]string{
		"limit":   limit_t,
		"country": country,
	}

	urlParams := c.prepareUrl(q)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	movies, err = unmarshalMovies(data)
	return
}

// InTheatersMovies, Retrieves movies currently in theaters.
func (c *Client) InTheatersMovies(page_limit int, page int, country string) (movies []Movie_, total int, err error) {

	endpoint := c.getEndpoint("InTheatersMovies", "")
	page_limit_t := strconv.Itoa(page_limit)
	page_t := strconv.Itoa(page)

	q := map[string]string{
		"page_limit": page_limit_t,
		"page":       page_t,
		"country":    country,
	}

	urlParams := c.prepareUrl(q)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	movies, total, err = unmarshalSearch(data)
	return
}

// OpeningMovies, Retrieves current opening movies.
func (c *Client) OpeningMovies(limit int, country string) (movies []Movie_, err error) {

	endpoint := c.getEndpoint("OpeningMovies", "")
	limit_t := strconv.Itoa(limit)

	q := map[string]string{
		"limit":   limit_t,
		"country": country,
	}

	urlParams := c.prepareUrl(q)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	movies, err = unmarshalMovies(data)
	return
}

// UpcomingMovies, Retrieves upcoming movies. Results are paginated if they go past the specified page limit
func (c *Client) UpcomingMovies(page_limit int, page int, country string) (movies []Movie_, total int, err error) {

	endpoint := c.getEndpoint("UpcomingMovies", "")
	page_limit_t := strconv.Itoa(page_limit)
	page_t := strconv.Itoa(page)

	q := map[string]string{
		"page_limit": page_limit_t,
		"page":       page_t,
		"country":    country,
	}

	urlParams := c.prepareUrl(q)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	movies, total, err = unmarshalSearch(data)
	return
}

// TopRentals, Retrieves the current top dvd rentals.
func (c *Client) TopRentals(limit int, country string) (movies []Movie_, err error) {

	endpoint := c.getEndpoint("TopRentals", "")
	limit_t := strconv.Itoa(limit)

	q := map[string]string{
		"limit":   limit_t,
		"country": country,
	}

	urlParams := c.prepareUrl(q)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	movies, err = unmarshalMovies(data)
	return
}

// CurrentReleaseDVDs, Retrieves current release dvds. Results are paginated if they go past the specified page limit.
func (c *Client) CurrentReleaseDVDs(page_limit int, page int, country string) (movies []Movie_, total int, err error) {

	endpoint := c.getEndpoint("CurrentReleaseDVDs", "")
	page_limit_t := strconv.Itoa(page_limit)
	page_t := strconv.Itoa(page)

	q := map[string]string{
		"page_limit": page_limit_t,
		"page":       page_t,
		"country":    country,
	}

	urlParams := c.prepareUrl(q)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	movies, total, err = unmarshalSearch(data)
	return
}

// NewReleaseDVDs, Retrieves new release dvds. Results are paginated if they go past the specified page limit.
func (c *Client) NewReleaseDVDs(page_limit int, page int, country string) (movies []Movie_, total int, err error) {

	endpoint := c.getEndpoint("NewReleaseDVDs", "")
	page_limit_t := strconv.Itoa(page_limit)
	page_t := strconv.Itoa(page)

	q := map[string]string{
		"page_limit": page_limit_t,
		"page":       page_t,
		"country":    country,
	}

	urlParams := c.prepareUrl(q)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	movies, total, err = unmarshalSearch(data)
	return
}

// UpcomingDVDs, Retrieves upcoming release dvds. Results are paginated if they go past the specified page limit.
func (c *Client) UpcomingDVDs(page_limit int, page int, country string) (movies []Movie_, total int, err error) {

	endpoint := c.getEndpoint("UpcomingDVDs", "")
	page_limit_t := strconv.Itoa(page_limit)
	page_t := strconv.Itoa(page)

	q := map[string]string{
		"page_limit": page_limit_t,
		"page":       page_t,
		"country":    country,
	}

	urlParams := c.prepareUrl(q)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	movies, total, err = unmarshalSearch(data)
	return
}
