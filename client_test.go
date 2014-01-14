package rottentomatoes

import "testing"

/*func TestApiKey(t *testing.T) {
	apikey := "123abc"
	c, _ := NewClient(apikey)

	if c.ApiKey != apikey {
		t.Errorf("NewClient(%v) = %v, want %v", apikey, c.ApiKey, apikey)
	}
}*/

func TestBaseUrls(t *testing.T) {
	c, _ := NewClient()

	if c.BaseUrl["MovieInfo"] != "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}.json?" {
		t.Errorf("wrong/missing BaseUrl value")
	}

	if c.BaseUrl["MovieCast"] != "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/cast.json?" {
		t.Errorf("wrong/missing BaseUrl value")
	}

	if c.BaseUrl["MovieClips"] != "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/clips.json?" {
		t.Errorf("wrong/missing BaseUrl value")
	}

	if c.BaseUrl["MovieReviews"] != "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/reviews.json?" {
		t.Errorf("wrong/missing BaseUrl value")
	}

	if c.BaseUrl["MovieSimilar"] != "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/similar.json?" {
		t.Errorf("wrong/missing BaseUrl value")
	}

	if c.BaseUrl["MovieAlias"] != "http://api.rottentomatoes.com/api/public/v1.0/movie_alias.json?" {
		t.Errorf("wrong/missing BaseUrl value")
	}

	if c.BaseUrl["MoviesSearch"] != "http://api.rottentomatoes.com/api/public/v1.0/movies.json?" {
		t.Errorf("wrong/missing BaseUrl value")
	}
}
