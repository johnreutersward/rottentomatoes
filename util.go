package rottentomatoes

import (
	"bytes"
	"encoding/json"
	"net/url"
	"text/template"
)

func (c *client) getEndpoint(tmplName string, id string) string {
	t, _ := template.New(tmplName).Parse(c.BaseUrl[tmplName])
	buf := new(bytes.Buffer)
	t.Execute(buf, id)
	return buf.String()
}

func (c *client) prepareUrl(p map[string]string) string {

	urlVals := url.Values{}
	urlVals.Set("apikey", c.ApiKey)

	for k, v := range p {
		urlVals.Set(k, v)
	}
	return urlVals.Encode()
}

func UnmarshalMoviesInfo(data []byte) (Movie, error) {
	var m Movie
	err := json.Unmarshal(data, &m)

	return m, err
}

func UnmarshalCastInfo(data []byte) ([]Cast, error) {
	var jsond map[string]*json.RawMessage
	_ = json.Unmarshal(data, &jsond)

	var castList []Cast
	err := json.Unmarshal(*jsond["cast"], &castList)

	return castList, err
}

func UnmarshalMovies(data []byte) ([]Movie_, error) {
	var jsond map[string]*json.RawMessage
	_ = json.Unmarshal(data, &jsond)

	var movieList []Movie_
	err := json.Unmarshal(*jsond["movies"], &movieList)

	return movieList, err
}

func UnmarshalSearch(data []byte) ([]Movie_, int, error) {
	var jsond map[string]*json.RawMessage
	_ = json.Unmarshal(data, &jsond)

	var movieList []Movie_
	err := json.Unmarshal(*jsond["movies"], &movieList)

	var t int
	err = json.Unmarshal(*jsond["total"], &t)

	return movieList, t, err
}

func UnmarshalReviews(data []byte) ([]Review, int, error) {
	var jsond map[string]*json.RawMessage
	_ = json.Unmarshal(data, &jsond)

	var reviewList []Review
	err := json.Unmarshal(*jsond["reviews"], &reviewList)

	var t int
	err = json.Unmarshal(*jsond["total"], &t)

	return reviewList, t, err
}

func UnmarshalClips(data []byte) ([]Clip, error) {
	var jsond map[string]*json.RawMessage
	_ = json.Unmarshal(data, &jsond)

	var clipList []Clip
	err := json.Unmarshal(*jsond["clips"], &clipList)

	return clipList, err
}
