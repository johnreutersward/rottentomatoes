# rotttentomatoes

[![travis-ci status](https://api.travis-ci.org/rojters/rottentomatoes.png)](https://travis-ci.org/rojters/rottentomatoes) [![GoDoc](https://godoc.org/github.com/rojters/rottentomatoes?status.png)](https://godoc.org/github.com/rojters/rottentomatoes) 

Rotten Tomatoes API wrapper for Go

## Usage

Start by applying for an API key at http://developer.rottentomatoes.com/ (it's free).

Store the API key in a enviromental variable called `ROTTENTOMATOES_APIKEY`.

Install by issuing: 

```bash
$ go get github.com/rojters/rottentomatoes
```

Example application:

```go
package main

import (
	"fmt"
	"github.com/rojters/rottentomatoes"
	"log"
)

func main() {
	rt, _ := rottentomatoes.NewClient()

	// Get info using movie id
	m, err := rt.MoviesInfo("14281")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Title: %s, Year: %d, Runtime: %d min\n", m.Title, m.Year, m.Runtime)
	// Title: The Big Lebowski, Year: 1998, Runtime: 118 min

}
```

## License

MIT
