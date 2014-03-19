package rottentomatoes

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetEndpoint(t *testing.T) {
	c, _ := NewClient(&http.Client{})

	got := c.getEndpoint("MovieInfo", "14281")
	want := "http://api.rottentomatoes.com/api/public/v1.0/movies/14281.json?"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("getEndpoint = %+v,\nwant %+v", got, want)
	}

	got = c.getEndpoint("MovieCast", "14281")
	want = "http://api.rottentomatoes.com/api/public/v1.0/movies/14281/cast.json?"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("getEndpoint = %+v,\nwant %+v", got, want)
	}

	got = c.getEndpoint("MovieClips", "14281")
	want = "http://api.rottentomatoes.com/api/public/v1.0/movies/14281/clips.json?"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("getEndpoint = %+v,\nwant %+v", got, want)
	}

	got = c.getEndpoint("MovieReviews", "14281")
	want = "http://api.rottentomatoes.com/api/public/v1.0/movies/14281/reviews.json?"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("getEndpoint = %+v,\nwant %+v", got, want)
	}

	got = c.getEndpoint("MovieSimilar", "14281")
	want = "http://api.rottentomatoes.com/api/public/v1.0/movies/14281/similar.json?"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("getEndpoint = %+v,\nwant %+v", got, want)
	}

	got = c.getEndpoint("MovieAlias", "")
	want = "http://api.rottentomatoes.com/api/public/v1.0/movie_alias.json?"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("getEndpoint = %+v,\nwant %+v", got, want)
	}

	got = c.getEndpoint("MoviesSearch", "")
	want = "http://api.rottentomatoes.com/api/public/v1.0/movies.json?"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("getEndpoint = %+v,\nwant %+v", got, want)
	}
}
