package rottentomatoes

import (
    "io/ioutil"
    "net/url"
    "bytes"
    "log"
    "net/http"
    "encoding/json"
    "text/template"
    "errors"
)

type Cast struct {
	Name string `json:"name"`
	Id string `json:"id"`
	Characters []string `json:"characters"`
}


func (c *Client) CastInfo(id string) ([]Cast, error) {

    if len(c.ApiKey) == 0  {
        return nil, errors.New("missing ApiKey")
    }

    t, _ := template.New("CastInfoUrl").Parse(c.BaseUrl["CastInfo"])
    buf := new(bytes.Buffer)
    t.Execute(buf, id)

    v := url.Values {}
    v.Set("apikey", c.ApiKey)

    endp := buf.String() + v.Encode()

    res, err := http.Get(endp)
    if err != nil {
        log.Fatal(err)
    }

    data, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        log.Fatal(err)
    }

    var jsond map[string]*json.RawMessage
    err = json.Unmarshal(data, &jsond)
    if err != nil {
        log.Fatal(err)
    }
    
    var castList []Cast
    errs := json.Unmarshal(*jsond["cast"], &castList)   
    if errs != nil {
        log.Fatal(errs)
    }

    return castList, errs
}



