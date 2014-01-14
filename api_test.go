package rottentomatoes

import (
	"reflect"
	"testing"
)

func TestMovieInfo(t *testing.T) {
	c, _ := NewClient()

	m, err := c.MovieInfo("14281")

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

	got := m.Title
	want := "The Big Lebowski"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MovieInfo = %+v,\nwant %+v", got, want)
	}
}

func TestMovieCast(t *testing.T) {
	c, _ := NewClient()

	cast, err := c.MovieCast("14281")

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

	got := cast[0].Name
	want := "Jeff Bridges"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MovieCast = %+v,\nwant %+v", got, want)
	}
}

func TestMovieClips(t *testing.T) {
	c, _ := NewClient()

	clips, err := c.MovieClips("14281")

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

	got := clips[0].Title
	want := "Big Lebowski Trailer"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MovieClips = %+v,\nwant %+v", got, want)
	}
}
