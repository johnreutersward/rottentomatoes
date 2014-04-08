# rottentomatoes

[![travis-ci status](https://api.travis-ci.org/rojters/rottentomatoes.png)](https://travis-ci.org/rojters/rottentomatoes) [![GoDoc](https://godoc.org/github.com/rojters/rottentomatoes?status.png)](https://godoc.org/github.com/rojters/rottentomatoes) 

rottentomatoes is a Go client library for accessing the Rotten Tomatoes API.

## Usage

Install: 

```bash
$ go get github.com/rojters/rottentomatoes
```

Example application:

```go
package main

import (
	"fmt"
	"github.com/rojters/rottentomatoes"
)

func main() {
	// Create a new client.
	rt := rottentomatoes.NewClient(nil, "")

	// Use IMDb Id to find movie information.
	// http://www.imdb.com/title/tt0118715/
	m, _ := rt.DetailedInfo.MovieAlias("0118715")

	// This will output: The Big Lebowski (1998), with: Jeff Bridges, John Goodman
	fmt.Printf("%s (%d), with: %s, %s", m.Title, m.Year, m.AbridgedCast[0].Name, m.AbridgedCast[1].Name)
}
```

## License

MIT
