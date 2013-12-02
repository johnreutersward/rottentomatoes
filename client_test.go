package rottentomatoes

import "testing"

func TestApiKey(t *testing.T) {
	apikey := "123abc"
	c := NewClient(apikey)

	if c.ApiKey != apikey {
		t.Errorf("NewClient(%v) = %v, want %v", apikey, c.ApiKey, apikey)
	}

}

func TestBaseUrls(t *testing.T) {
	apikey := "123abc"
	c := NewClient(apikey)
	if c.BaseUrl["CastInfo"] != "http://api.rottentomatoes.com/api/public/v1.0/movies/{{.}}/cast.json?" {
		t.Errorf("wrong BaseUrl value")
	}
}