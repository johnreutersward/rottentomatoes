package rottentomatoes

import (
	"log"
	"reflect"
	"testing"
)

func TestUnmarshalMoviesInfo(t *testing.T) {
	data := `{
	    "id": 14281,
    	"title": "The Big Lebowski",
    	"year": 1998,
		"genres": ["Comedy"],
	    "mpaa_rating": "R",
	    "runtime": 118,
	    "critics_consensus": "",
	    "release_dates": {
	        "theater": "1998-03-06",
	        "dvd": "1998-10-27"
    	},
    	"ratings": {
	        "critics_rating": "Certified Fresh",
	        "critics_score": 80,
	        "audience_rating": "Upright",
	        "audience_score": 94
    	}
	}`

	got, err := UnmarshalMoviesInfo([]byte(data))

	if err != nil {
		log.Panic(err)
	}

	want := Movie{
		Id:               14281,
		Title:            "The Big Lebowski",
		Year:             1998,
		Genres:           []string{"Comedy"},
		MPAA_Rating:      "R",
		Runtime:          118,
		CriticsConsensus: "",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UnmarshalMoviesInfo = %+v,\nwant %+v", got, want)
	}

}
