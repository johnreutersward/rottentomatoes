package rottentomatoes

import (
    "net/url"
    "bytes"
    "encoding/json"
    "text/template"
    "errors"
)

type Cast struct {
	Name string `json:"name"`
	Id string `json:"id"`
	Characters []string `json:"characters"`
}


func UnmarshalCastInfo(data []byte) ([]Cast, error) {
    var jsond map[string]*json.RawMessage
    err := json.Unmarshal(data, &jsond)
    if err != nil {
        //log.Fatal(err)
    }
    var castList []Cast
    err = json.Unmarshal(*jsond["cast"], &castList)
    if err != nil {
        //log.Fatal(errs)
    }
    return castList, err
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

    data, err := c.Request(endp)
    
    if err != nil {
        return nil, err
    }
    
    castList, err := UnmarshalCastInfo(data) 

    return castList, err
}



