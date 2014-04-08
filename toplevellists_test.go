package rottentomatoes

import (
	"testing"
)

func TestListsDirectory(t *testing.T) {

	rt := NewClient(nil, "")

	_, err := rt.TopLevelLists.ListsDirectory()

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestMovieListsDirectory(t *testing.T) {

	rt := NewClient(nil, "")

	_, err := rt.TopLevelLists.MovieListsDirectory()

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestDVDListsDirectory(t *testing.T) {

	rt := NewClient(nil, "")

	_, err := rt.TopLevelLists.DVDListsDirectory()

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}
