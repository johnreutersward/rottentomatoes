package rottentomatoes

import (
	"testing"
)

func TestBoxOffice(t *testing.T) {

	rt := NewClient(nil, "")

	opt := &Options{Limit: 1}

	_, err := rt.MovieLists.BoxOffice(opt)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestInTheaters(t *testing.T) {

	rt := NewClient(nil, "")

	_, err := rt.MovieLists.InTheaters(nil)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestOpeningMovies(t *testing.T) {

	rt := NewClient(nil, "")

	opt := &Options{Limit: 1}

	_, err := rt.MovieLists.OpeningMovies(opt)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestUpcomingMovies(t *testing.T) {

	rt := NewClient(nil, "")

	opt := &Options{PageLimit: 1}

	_, err := rt.MovieLists.UpcomingMovies(opt)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}
