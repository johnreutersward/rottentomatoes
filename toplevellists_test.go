package rottentomatoes

import (
	"testing"
	"time"
)

// There is a 1 second delay for each test,
// this is because the Rotten Tomatoes API enforces
// a hard limit on the number of simultaneous connections.

func TestListsDirectory(t *testing.T) {
	time.Sleep(1 * time.Second)

	rt := NewClient(nil, "")

	_, err := rt.TopLevelLists.ListsDirectory()

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestMovieListsDirectory(t *testing.T) {
	time.Sleep(1 * time.Second)

	rt := NewClient(nil, "")

	_, err := rt.TopLevelLists.MovieListsDirectory()

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestDVDListsDirectory(t *testing.T) {
	time.Sleep(1 * time.Second)

	rt := NewClient(nil, "")

	_, err := rt.TopLevelLists.DVDListsDirectory()

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}
