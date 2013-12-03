package rottentomatoes

import "testing"

func TestUnmarshalCastInfo(t *testing.T) {
	data := `{
	    "cast": [{
	        "id": "162655890",
	        "name": "Jeff Bridges",
	        "characters": ["The Dude"]
	    }, {
	        "id": "162655706",
	        "name": "John Goodman",
	        "characters": ["Walter Sobchak"]
	    }, {
	        "id": "162654248",
	        "name": "Julianne Moore",
	        "characters": ["Maude Lebowski"]
	    }, {
	        "id": "162652875",
	        "name": "Steve Buscemi",
	        "characters": ["Donny"]
	    }],
	    "links": {
	        "rel": "http://api.rottentomatoes.com/api/public/v1.0/movies/14281.json"
	    }
	}`
	result, _ := UnmarshalCastInfo([]byte(data))
	if result == nil {
		t.Errorf("result was empty")
	}

	c1 := &Cast{Name: "Jeff Bridges", Id: "162655890", Characters: []string{"The Dude"},}
	c1 := &Cast{Name: "John Goodman", Id: "162655706", Characters: []string{"Walter Sobchak"},}
	c1 := &Cast{Name: "Julianne Moore", Id: "162654248", Characters: []string{"Maude Lebowski"},}
	c1 := &Cast{Name: "Steve Buscemi", Id: "162652875", Characters: []string{"Donny"},}

	true := []Cast{}

}