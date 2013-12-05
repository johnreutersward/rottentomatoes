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
    	},
	    "synopsis": "",
	    "posters": {
	        "thumbnail": "http://content6.flixster.com/movie/10/95/21/10952108_mob.jpg",
	        "profile": "http://content6.flixster.com/movie/10/95/21/10952108_pro.jpg",
	        "detailed": "http://content6.flixster.com/movie/10/95/21/10952108_det.jpg",
	        "original": "http://content6.flixster.com/movie/10/95/21/10952108_ori.jpg"
	    },
	    "abridged_cast": [{
	        "name": "Jeff Bridges",
	        "id": "162655890",
	        "characters": ["The Dude"]
	    }, {
	        "name": "John Goodman",
	        "id": "162655706",
	        "characters": ["Walter Sobchak"]
	    }, {
	        "name": "Julianne Moore",
	        "id": "162654248",
	        "characters": ["Maude Lebowski"]
	    }, {
	        "name": "Steve Buscemi",
	        "id": "162652875",
	        "characters": ["Donny"]
	    }, {
	        "name": "David Huddleston",
	        "id": "167963972",
	        "characters": ["The Big Lebowski"]
	    }],
	    "abridged_directors": [{
	        "name": "Joel Coen"
	    }, {
	        "name": "Ethan Coen"
	    }],
	    "studio": "Gramercy Pictures",
	    "alternate_ids": {
	        "imdb": "0118715"
	    },
	    "links": {
	        "self": "http://api.rottentomatoes.com/api/public/v1.0/movies/14281.json",
	        "alternate": "http://www.rottentomatoes.com/m/big_lebowski/",
	        "cast": "http://api.rottentomatoes.com/api/public/v1.0/movies/14281/cast.json",
	        "clips": "http://api.rottentomatoes.com/api/public/v1.0/movies/14281/clips.json",
	        "reviews": "http://api.rottentomatoes.com/api/public/v1.0/movies/14281/reviews.json",
	        "similar": "http://api.rottentomatoes.com/api/public/v1.0/movies/14281/similar.json"
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
		ReleaseDates: ReleaseDates{
			Theater: "1998-03-06",
			DVD:     "1998-10-27",
		},
		Ratings: Ratings{
			CriticsRating:  "Certified Fresh",
			CriticsScore:   80,
			AudienceRating: "Upright",
			AudienceScore:  94,
		},
		Synopsis: "",
		Posters: Posters{
			Thumbnail: "http://content6.flixster.com/movie/10/95/21/10952108_mob.jpg",
			Profile:   "http://content6.flixster.com/movie/10/95/21/10952108_pro.jpg",
			Detailed:  "http://content6.flixster.com/movie/10/95/21/10952108_det.jpg",
			Original:  "http://content6.flixster.com/movie/10/95/21/10952108_ori.jpg",
		},
		AbridgedCast: []Cast{
			Cast{Name: "Jeff Bridges", Id: "162655890", Characters: []string{"The Dude"}},
			Cast{Name: "John Goodman", Id: "162655706", Characters: []string{"Walter Sobchak"}},
			Cast{Name: "Julianne Moore", Id: "162654248", Characters: []string{"Maude Lebowski"}},
			Cast{Name: "Steve Buscemi", Id: "162652875", Characters: []string{"Donny"}},
			Cast{Name: "David Huddleston", Id: "167963972", Characters: []string{"The Big Lebowski"}},
		},
		AbridgedDirectors: []Director{
			Director{Name: "Joel Coen"},
			Director{Name: "Ethan Coen"},
		},
		Studio: "Gramercy Pictures",
		AlternateIds: AlternateIds{
			IMDB: "0118715",
		},
		Links: Links{
			Self:      "http://api.rottentomatoes.com/api/public/v1.0/movies/14281.json",
			Alternate: "http://www.rottentomatoes.com/m/big_lebowski/",
			Cast:      "http://api.rottentomatoes.com/api/public/v1.0/movies/14281/cast.json",
			Clips:     "http://api.rottentomatoes.com/api/public/v1.0/movies/14281/clips.json",
			Reviews:   "http://api.rottentomatoes.com/api/public/v1.0/movies/14281/reviews.json",
			Similar:   "http://api.rottentomatoes.com/api/public/v1.0/movies/14281/similar.json",
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UnmarshalMoviesInfo = %+v,\nwant %+v", got, want)
	}

}

func TestUnmarshalMovies(t *testing.T) {
	data := `{
		"movies": [{
			    "id": "104762093",
	        	"title": "Fargo",
	        	"year": 1996
			}, {
				"id": "12776",
	        	"title": "O Brother, Where Art Thou?",
	        	"year": 2000
			}]
	}`

	got, err := UnmarshalMovies([]byte(data))

	if err != nil {
		log.Panic(err)
	}

	want := []Movies{
		Movies{Id: "104762093", Title: "Fargo", Year: 1996},
		Movies{Id: "12776", Title: "O Brother, Where Art Thou?", Year: 2000},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UnmarshalMovies = %+v,\nwant %+v", got, want)
	}

}

func TestUnmarshalSearch(t *testing.T) {
	data := `{
		"total": 2,
		"movies": [{
			    "id": "104762093",
	        	"title": "Fargo",
	        	"year": 1996
			}, {
				"id": "770717155",
        		"title": "Fargo Express",
        		"year": 1932
			}]
	}`

	got, total, err := UnmarshalSearch([]byte(data))

	if err != nil {
		log.Panic(err)
	}

	want_total := 2
	want := []Movies{
		Movies{Id: "104762093", Title: "Fargo", Year: 1996},
		Movies{Id: "770717155", Title: "Fargo Express", Year: 1932},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UnmarshalSearch = %+v,\nwant %+v", got, want)
	}

	if total != want_total {
		t.Errorf("UnmarshalSearch/total = %+v,\nwant %+v", total, want_total)
	}

}

func TestUnmarshalReviews(t *testing.T) {
	data := `{
		"total": 2,
    	"reviews": [{
	        "critic": "Eric Hynes",
	        "date": "2013-08-04",
	        "freshness": "fresh",
	        "publication": "Village Voice",
	        "quote": "When teenaged Andy plops down on the grass to share his old toys with a shy little girl, the film spikes with sadness and layered pleasure -- a concise, deeply wise expression of the ephemeral that feels real and yet utterly transporting.",
	        "links": {
	            "review": "http://www.villagevoice.com/2010-06-15/film/toys-are-us-in-toy-story-3/full/"
	        }
        }, {
	        "critic": "David Denby",
	        "date": "2013-08-04",
	        "freshness": "fresh",
	        "publication": "New Yorker",
	        "quote": "There are many sweet laughs and joking allusions to horror and prison-break movies, but the Pixar gang gets at the most primary fear -- being cast off and no longer of use.",
	        "links": {
	            "review": "http://www.newyorker.com/arts/reviews/film/toy_story_3_unkrich"
	        }
        }]
    }`

	got, total, err := UnmarshalReviews([]byte(data))

	if err != nil {
		log.Panic(err)
	}

	want_total := 2

	want := []Review{
		Review{
			Critic:      "Eric Hynes",
			Date:        "2013-08-04",
			Freshness:   "fresh",
			Publication: "Village Voice",
			Quote:       "When teenaged Andy plops down on the grass to share his old toys with a shy little girl, the film spikes with sadness and layered pleasure -- a concise, deeply wise expression of the ephemeral that feels real and yet utterly transporting.",
			Links:       ReviewLinks{Review: "http://www.villagevoice.com/2010-06-15/film/toys-are-us-in-toy-story-3/full/"},
		},
		Review{
			Critic:      "David Denby",
			Date:        "2013-08-04",
			Freshness:   "fresh",
			Publication: "New Yorker",
			Quote:       "There are many sweet laughs and joking allusions to horror and prison-break movies, but the Pixar gang gets at the most primary fear -- being cast off and no longer of use.",
			Links:       ReviewLinks{Review: "http://www.newyorker.com/arts/reviews/film/toy_story_3_unkrich"},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UnmarshalReviews = %+v,\nwant %+v", got, want)
	}

	if total != want_total {
		t.Errorf("UnmarshalReviews/total = %+v,\nwant %+v", total, want_total)
	}

}

func TestUnmarshalClips(t *testing.T) {
	data := `{
		"clips": [{
	        "title": "Toy Story 3",
	        "duration": "149",
	        "thumbnail": "http://content.internetvideoarchive.com/content/photos/3465/14553712_.jpg",
	        "links": {
	            "alternate": "http://www.rottentomatoes.com/m/toy_story_3/trailers/11028566"
	        }
    	}, {
	        "title": "Toy Story 3: Meet Ken",
	        "duration": "102",
	        "thumbnail": "http://content.internetvideoarchive.com/content/photos/6756/28378422_.jpg",
	        "links": {
	            "alternate": "http://www.rottentomatoes.com/m/toy_story_3/trailers/11108601"
	        }
    	}]
    }`

	got, err := UnmarshalClips([]byte(data))

	if err != nil {
		log.Panic(err)
	}

	want := []Clip{
		Clip{
			Title:     "Toy Story 3",
			Duration:  "149",
			Thumbnail: "http://content.internetvideoarchive.com/content/photos/3465/14553712_.jpg",
			Links:     ClipLinks{Alternate: "http://www.rottentomatoes.com/m/toy_story_3/trailers/11028566"},
		},
		Clip{
			Title:     "Toy Story 3: Meet Ken",
			Duration:  "102",
			Thumbnail: "http://content.internetvideoarchive.com/content/photos/6756/28378422_.jpg",
			Links:     ClipLinks{Alternate: "http://www.rottentomatoes.com/m/toy_story_3/trailers/11108601"},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UnmarshalClips = %+v,\nwant %+v", got, want)
	}

}
