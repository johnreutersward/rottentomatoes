package rottentomatoes

import (
	"testing"
)

func TestTopRentals(t *testing.T) {

	rt := NewClient(nil, "")

	opt := &Options{Limit: 1, PageLimit: 1}

	_, err := rt.DVDLists.TopRentals(opt)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestCurrentReleaseDVDs(t *testing.T) {

	rt := NewClient(nil, "")

	opt := &Options{Limit: 1, PageLimit: 1}

	_, err := rt.DVDLists.CurrentReleaseDVDs(opt)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestNewReleaseDVDs(t *testing.T) {

	rt := NewClient(nil, "")

	opt := &Options{Limit: 1, PageLimit: 1}

	_, err := rt.DVDLists.NewReleaseDVDs(opt)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}

func TestUpcomingDVDs(t *testing.T) {

	rt := NewClient(nil, "")

	opt := &Options{Limit: 1, PageLimit: 1}

	_, err := rt.DVDLists.UpcomingDVDs(opt)

	if err != nil {
		t.Errorf("err not nil: ", err)
	}

}
