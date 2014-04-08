package rottentomatoes

type SearchService struct {
	client *Client
}

// Returns movies from search results.
// Available options: PageLimit, Page.
func (s *SearchService) MovieSearch(query string, opt *Options) (*MovieListsResponse, error) {
	rel := "movies.json?"
	if opt == nil {
		opt = &Options{}
	}
	opt.Query = query
	resp := new(MovieListsResponse)
	_, err := s.client.get(rel, opt, resp)
	return resp, err
}
