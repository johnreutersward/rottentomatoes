package rottentomatoes

import (
    "net/url"
    "bytes"
    //"log"
    "encoding/json"
    //"fmt"
    "text/template"
    "errors"
)

type Error struct {
    Error string `json:"error"`
}

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

    data, err := c.Request(endp)
    
    if err != nil {
        return nil, err
    }

    //var f interface{}
    //_ = json.Unmarshal(data, &f)


    var jsond map[string]*json.RawMessage
    err = json.Unmarshal(data, &jsond)
    if err != nil {
        //log.Fatal(err)
    }
    var castList []Cast
    errs := json.Unmarshal(*jsond["cast"], &castList)
    if errs != nil {
        //log.Fatal(errs)
    }

    //fmt.Printf("%+v", castList)
    //return castList, errs

    // m := f.(map[string]interface{})
    // b := []byte(m["cast"])

    // var castList []Cast
    // err = json.Unmarshal(b, &castList)   

    return castList, err
}



