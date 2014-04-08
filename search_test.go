package rottentomatoes

import (
	"testing"
)

func TestMovieSearch(t *testing.T) {

	rt := NewClient(nil, "")

	_, err := rt.Search.MovieSearch("The Big Lebowski", nil)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}
