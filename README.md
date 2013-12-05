# rotttentomatoes

Rotten Tomatoes API wrapper for Go

## Usage

Start by applying for an API key at http://developer.rottentomatoes.com/ (it's free).
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
	rt := rottentomatoes.NewClient("YOUR_API_KEY")

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
