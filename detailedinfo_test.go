package rottentomatoes

import (
	"testing"
	"time"
)

// There is a 1 second delay for each test,
// this is because the Rotten Tomatoes API enforces
// a hard limit on the number of simultaneous connections.

func TestMoviesInfo(t *testing.T) {
	time.Sleep(1 * time.Second)

	rt := NewClient(nil, "")

	_, err := rt.DetailedInfo.MovieInfo("770672122")

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestCastInfo(t *testing.T) {
	time.Sleep(1 * time.Second)

	rt := NewClient(nil, "")

	_, err := rt.DetailedInfo.CastInfo("770672122")

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestMovieClips(t *testing.T) {
	time.Sleep(1 * time.Second)

	rt := NewClient(nil, "")

	_, err := rt.DetailedInfo.MovieClips("770672122")

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestMovieReviews(t *testing.T) {
	time.Sleep(1 * time.Second)

	rt := NewClient(nil, "")

	o := &Options{PageLimit: 1}

	_, err := rt.DetailedInfo.MovieReviews("770672122", o)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestMovieSimilar(t *testing.T) {
	time.Sleep(1 * time.Second)

	rt := NewClient(nil, "")

	o := &Options{Limit: 1}

	_, err := rt.DetailedInfo.MovieSimilar("770672122", o)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestMovieAlias(t *testing.T) {
	time.Sleep(1 * time.Second)

	rt := NewClient(nil, "")

	_, err := rt.DetailedInfo.MovieAlias("0031381")

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}
