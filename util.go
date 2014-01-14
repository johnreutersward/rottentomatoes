package rottentomatoes

import (
	"bytes"
	"encoding/json"
	"net/url"
	"text/template"
)

func (c *Client) getEndpoint(tmplName string, id string) string {
	t, _ := template.New(tmplName).Parse(c.BaseUrl[tmplName])
	buf := new(bytes.Buffer)
	t.Execute(buf, id)
	return buf.String()
}

func (c *Client) prepareUrl(p map[string]string) string {

	urlVals := url.Values{}
	urlVals.Set("apikey", c.ApiKey)

	for k, v := range p {
		urlVals.Set(k, v)
	}
	return urlVals.Encode()
}

func unmarshalMoviesInfo(data []byte) (Movie, error) {
	var m Movie
	err := json.Unmarshal(data, &m)

	return m, err
}

func unmarshalCastInfo(data []byte) ([]Cast, error) {
	var jsond map[string]*json.RawMessage
	_ = json.Unmarshal(data, &jsond)

	var castList []Cast
	err := json.Unmarshal(*jsond["cast"], &castList)

	return castList, err
}

func unmarshalMovies(data []byte) ([]Movie_, error) {
	var jsond map[string]*json.RawMessage
	_ = json.Unmarshal(data, &jsond)

	var movieList []Movie_
	err := json.Unmarshal(*jsond["movies"], &movieList)

	return movieList, err
}

func unmarshalSearch(data []byte) ([]Movie_, int, error) {
	var jsond map[string]*json.RawMessage
	_ = json.Unmarshal(data, &jsond)

	var movieList []Movie_
	err := json.Unmarshal(*jsond["movies"], &movieList)

	var t int
	err = json.Unmarshal(*jsond["total"], &t)

	return movieList, t, err
}

func unmarshalReviews(data []byte) ([]Review, int, error) {
	var jsond map[string]*json.RawMessage
	_ = json.Unmarshal(data, &jsond)

	var reviewList []Review
	err := json.Unmarshal(*jsond["reviews"], &reviewList)

	var t int
	err = json.Unmarshal(*jsond["total"], &t)

	return reviewList, t, err
}

func unmarshalClips(data []byte) ([]Clip, error) {
	var jsond map[string]*json.RawMessage
	_ = json.Unmarshal(data, &jsond)

	var clipList []Clip
	err := json.Unmarshal(*jsond["clips"], &clipList)

	return clipList, err
}
