package rottentomatoes

import (
	"testing"
	"time"
)

// There is a 1 second delay for each test,
// this is because the Rotten Tomatoes API enforces
// a hard limit on the number of simultaneous connections.

func TestTopRentals(t *testing.T) {
	time.Sleep(1 * time.Second)

	rt := NewClient(nil, "")

	opt := &Options{Limit: 1, PageLimit: 1}

	_, err := rt.DVDLists.TopRentals(opt)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestCurrentReleaseDVDs(t *testing.T) {
	time.Sleep(1 * time.Second)

	rt := NewClient(nil, "")

	opt := &Options{Limit: 1, PageLimit: 1}

	_, err := rt.DVDLists.CurrentReleaseDVDs(opt)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestNewReleaseDVDs(t *testing.T) {
	time.Sleep(1 * time.Second)

	rt := NewClient(nil, "")

	opt := &Options{Limit: 1, PageLimit: 1}

	_, err := rt.DVDLists.NewReleaseDVDs(opt)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestUpcomingDVDs(t *testing.T) {
	time.Sleep(1 * time.Second)

	rt := NewClient(nil, "")

	opt := &Options{Limit: 1, PageLimit: 1}

	_, err := rt.DVDLists.UpcomingDVDs(opt)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}
