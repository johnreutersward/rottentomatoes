# [![Rotten Tomatoes API](http://i.imgur.com/1XiPsZY.png?2)](http:rottentomatoes.com/) rottentomatoes #

rottentomatoes is a Go client library for accessing the Rotten Tomatoes API.

**Documentation:** <http://godoc.org/github.com/rojters/rottentomatoes>
**Build Status:** [![travis-ci status](https://api.travis-ci.org/rojters/rottentomatoes.png)](https://travis-ci.org/rojters/rottentomatoes)
**Rotten Tomatoes API Documentation:** <http://developer.rottentomatoes.com/>

## Usage ##

```go
import "github.com/rojters/rottentomatoes"
```

Construct a new Rotten Tomatoes client, then use the various services on the client to
access different parts of the Rotten Tomatoes API. For example, to find information
about a particular movie using it's IMDb Id:

```go
rt := rottentomatoes.NewClient(nil, "")
// http://www.imdb.com/title/tt0118715/ (The Big Lebowski)
m, _ := rt.DetailedInfo.MovieAlias("0118715")
```

Some API methods have optional parameters that can be passed. For example, 
to page limit the number of results returned:

```go
rt := rottentomatoes.NewClient(nil, "")
opt := &rottentomatoes.Options{PageLimit: 10}
it, _ := rt.MovieLists.InTheaters(Opt)
```

## License ##

MIT
