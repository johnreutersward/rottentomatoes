package rottentomatoes

import (
	"testing"
)

func TestGetEndpoint(t *testing.T) {
	c, _ := NewClient()

	got := c.getEndpoint("MovieInfo", "14281")
	expected := "http://api.rottentomatoes.com/api/public/v1.0/movies/14281.json?"
	if got != expected {
		t.Errorf("incorrect endpoint encode")
	}

	got = c.getEndpoint("MovieCast", "14281")
	expected = "http://api.rottentomatoes.com/api/public/v1.0/movies/14281/cast.json?"
	if got != expected {
		t.Errorf("incorrect endpoint encode")
	}

	got = c.getEndpoint("MovieClips", "14281")
	expected = "http://api.rottentomatoes.com/api/public/v1.0/movies/14281/clips.json?"
	if got != expected {
		t.Errorf("incorrect endpoint encode")
	}

	got = c.getEndpoint("MovieReviews", "14281")
	expected = "http://api.rottentomatoes.com/api/public/v1.0/movies/14281/reviews.json?"
	if got != expected {
		t.Errorf("incorrect endpoint encode")
	}

	got = c.getEndpoint("MovieSimilar", "14281")
	expected = "http://api.rottentomatoes.com/api/public/v1.0/movies/14281/similar.json?"
	if got != expected {
		t.Errorf("incorrect endpoint encode")
	}

	got = c.getEndpoint("MovieAlias", "")
	expected = "http://api.rottentomatoes.com/api/public/v1.0/movie_alias.json?"
	if got != expected {
		t.Errorf("incorrect endpoint encode")
	}

	got = c.getEndpoint("MoviesSearch", "")
	expected = "http://api.rottentomatoes.com/api/public/v1.0/movies.json?"
	if got != expected {
		t.Errorf("incorrect endpoint encode")
	}

}
