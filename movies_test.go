package rottentomatoes

import (
	"testing"
	"reflect"
)

func TestUnmarshalMoviesInfo(t *testing.T) {
	data := `{
	    "id": 14281,
    	"title": "The Big Lebowski"
	}`
	
	got, _ := UnmarshalMoviesInfo([]byte(data))

	want := Movie {
		Id: 14281,
		Title: "The Big Lebowski",
	}

	if !reflect.DeepEqual(got, want) {
    	t.Errorf("UnmarshalMoviesInfo = %+v, want %+v", got, want)
    }


}