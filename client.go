package rottentomatoes

type Client struct {
	ApiKey string
	BaseUrl map[string]string
}

func NewClient(apikey string) *Client {
	c := &Client {
		ApiKey: apikey,
		BaseUrl:  map[string]string {
			"MoviesInfo": "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}.json?",
			"CastInfo": "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/cast.json?",
			"MovieClips": "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/clips.json?",
			"MovieReviews": "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/reviews.json?",
			"MoviesSimilar": "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/similar.json?",
			"MoviesAlias": "http://api.rottentomatoes.com/api/public/v1.0/movie_alias.json?",
			"Search": "http://api.rottentomatoes.com/api/public/v1.0/movies.json?",

		},
	}

	return c
}