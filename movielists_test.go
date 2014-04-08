package rottentomatoes

import (
	"testing"
	"time"
)

// There is a 1 second delay for each test,
// this is because the Rotten Tomatoes API enforces
// a hard limit on the number of simultaneous connections.

func TestBoxOffice(t *testing.T) {
	time.Sleep(1 * time.Second)

	rt := NewClient(nil, "")

	opt := &Options{Limit: 1}

	_, err := rt.MovieLists.BoxOffice(opt)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestInTheaters(t *testing.T) {
	time.Sleep(1 * time.Second)

	rt := NewClient(nil, "")

	_, err := rt.MovieLists.InTheaters(nil)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestOpeningMovies(t *testing.T) {
	time.Sleep(1 * time.Second)

	rt := NewClient(nil, "")

	opt := &Options{Limit: 1}

	_, err := rt.MovieLists.OpeningMovies(opt)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestUpcomingMovies(t *testing.T) {
	time.Sleep(1 * time.Second)

	rt := NewClient(nil, "")

	opt := &Options{PageLimit: 1}

	_, err := rt.MovieLists.UpcomingMovies(opt)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}
