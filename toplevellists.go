package rottentomatoes

type ListsDirectoryResponse struct {
	Links ListsDirectoryLinks `json:"links,omitempty"`
}

type ListsDirectoryLinks struct {
	Movies string `json:"movies,omitempty"`
	DVDs   string `json:"dvds,omitempty"`
}

type MovieListsDirectoryResponse struct {
	Links MovieListsDirectoryLinks `json:"links,omitempty"`
}

type MovieListsDirectoryLinks struct {
	BoxOffice  string `json:"box_office,omitempty"`
	InTheaters string `json:"in_theaters,omitempty"`
	Opening    string `json:"opening,omitempty"`
	Upcoming   string `json:"upcoming,omitempty"`
}

type DVDListsDirectoryResponse struct {
	Links DVDListsDirectoryLinks `json:"links,omitempty"`
}

type DVDListsDirectoryLinks struct {
	TopRentals      string `json:"top_rentals,omitempty"`
	CurrentReleases string `json:"current_releases,omitempty"`
	NewReleases     string `json:"new_releases,omitempty"`
	Upcoming        string `json:"upcoming,omitempty"`
}

type TopLevelListsService struct {
	client *Client
}

// Returns the top level lists available in the API.
func (t *TopLevelListsService) ListsDirectory() (*ListsDirectoryResponse, error) {
	rel := "lists.json?"
	resp := new(ListsDirectoryResponse)
	_, err := t.client.get(rel, nil, resp)
	return resp, err
}

// Returns the movie lists.
func (t *TopLevelListsService) MovieListsDirectory() (*MovieListsDirectoryResponse, error) {
	rel := "lists/movies.json?"
	resp := new(MovieListsDirectoryResponse)
	_, err := t.client.get(rel, nil, resp)
	return resp, err
}

// Returns the dvd lists.
func (t *TopLevelListsService) DVDListsDirectory() (*DVDListsDirectoryResponse, error) {
	rel := "lists/dvds.json?"
	resp := new(DVDListsDirectoryResponse)
	_, err := t.client.get(rel, nil, resp)
	return resp, err
}
