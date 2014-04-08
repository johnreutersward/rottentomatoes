package rottentomatoes

type DVDListsService struct {
	client *Client
}

// Returns the current top dvd rentals.
// Available options: Limit, Country.
func (d *DVDListsService) TopRentals(opt *Options) (*MovieListsResponse, error) {
	rel := "lists/dvds/top_rentals.json?"
	resp := new(MovieListsResponse)
	_, err := d.client.get(rel, opt, resp)
	return resp, err
}

// Returns current release dvds. Results are paginated if they go past the specified page limit.
// Available options: PageLimit, Page, Country.
func (d *DVDListsService) CurrentReleaseDVDs(opt *Options) (*MovieListsResponse, error) {
	rel := "lists/dvds/current_releases.json?"
	resp := new(MovieListsResponse)
	_, err := d.client.get(rel, opt, resp)
	return resp, err
}

// Returns new release dvds. Results are paginated if they go past the specified page limit.
// Available options: PageLimit, Page, Country.
func (d *DVDListsService) NewReleaseDVDs(opt *Options) (*MovieListsResponse, error) {
	rel := "lists/dvds/new_releases.json?"
	resp := new(MovieListsResponse)
	_, err := d.client.get(rel, opt, resp)
	return resp, err
}

// Returns upcoming dvds. Results are paginated if they go past the specified page limit.
// Available options: PageLimit, Page, Country.
func (d *DVDListsService) UpcomingDVDs(opt *Options) (*MovieListsResponse, error) {
	rel := "lists/dvds/upcoming.json?"
	resp := new(MovieListsResponse)
	_, err := d.client.get(rel, opt, resp)
	return resp, err
}
