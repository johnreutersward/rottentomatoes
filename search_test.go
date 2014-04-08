package rottentomatoes

import (
	"testing"
	"time"
)

// There is a 1 second delay for each test,
// this is because the Rotten Tomatoes API enforces
// a hard limit on the number of simultaneous connections.

func TestMovieSearch(t *testing.T) {
	time.Sleep(1 * time.Second)

	rt := NewClient(nil, "")

	_, err := rt.Search.MovieSearch("The Big Lebowski", nil)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}
