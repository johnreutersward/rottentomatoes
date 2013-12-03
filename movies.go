package rottentomatoes

import (
    "net/url"
    "bytes"
    "encoding/json"
    "text/template"
    "errors"
)

type Movie struct {
    Id int `json:"id"`
    Title string `json:"title"`
}


func UnmarshalMoviesInfo(data []byte) (Movie, error) {
    var m Movie
    //m := new(Movie)
    err := json.Unmarshal(data, &m)

    return m, err
}


func (c *Client) MoviesInfo(id string) (Movie, error) {

    var movie Movie

    if len(c.ApiKey) == 0  {
        return movie, errors.New("missing ApiKey")
    }

    t, _ := template.New("MoviesInfoUrl").Parse(c.BaseUrl["MoviesInfo"])
    buf := new(bytes.Buffer)
    t.Execute(buf, id)

    v := url.Values {}
    v.Set("apikey", c.ApiKey)

    endp := buf.String() + v.Encode()

    data, err := c.Request(endp)
    
    if err != nil {
        return movie, err
    }
    
    movie, err = UnmarshalMoviesInfo(data) 

    return movie, err
}



