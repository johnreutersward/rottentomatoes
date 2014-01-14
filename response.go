package rottentomatoes

type Cast struct {
	Name       string   `json:"name"`
	Id         string   `json:"id"`
	Characters []string `json:"characters"`
}

type Movie struct {
	Id                int          `json:"id"`
	Title             string       `json:"title"`
	Year              int          `json:"year"`
	Genres            []string     `json:"genres"`
	MPAA_Rating       string       `json:"mpaa_rating"`
	Runtime           int          `json:"runtime"`
	CriticsConsensus  string       `json:"critics_consensus"`
	ReleaseDates      ReleaseDates `json:"release_dates"`
	Ratings           Ratings      `json:"ratings"`
	Synopsis          string       `json:"synopsis"`
	Posters           Posters      `json:"posters"`
	AbridgedCast      []Cast       `json:"abridged_cast"`
	AbridgedDirectors []Director   `json:"abridged_directors"`
	Studio            string       `json:"studio"`
	AlternateIds      AlternateIds `json:"alternate_ids"`
	Links             Links        `json:"links"`
}

// Same as Movie but Id is of type string.
type Movie_ struct {
	Id                string       `json:"id"`
	Title             string       `json:"title"`
	Year              int          `json:"year"`
	Genres            []string     `json:"genres"`
	MPAA_Rating       string       `json:"mpaa_rating"`
	Runtime           int          `json:"runtime"`
	CriticsConsensus  string       `json:"critics_consensus"`
	ReleaseDates      ReleaseDates `json:"release_dates"`
	Ratings           Ratings      `json:"ratings"`
	Synopsis          string       `json:"synopsis"`
	Posters           Posters      `json:"posters"`
	AbridgedCast      []Cast       `json:"abridged_cast"`
	AbridgedDirectors []Director   `json:"abridged_directors"`
	Studio            string       `json:"studio"`
	AlternateIds      AlternateIds `json:"alternate_ids"`
	Links             Links        `json:"links"`
}

type ReleaseDates struct {
	Theater string `json:"theater"`
	DVD     string `json:"dvd"`
}

type Ratings struct {
	CriticsRating  string `json:"critics_rating"`
	CriticsScore   int    `json:"critics_score"`
	AudienceRating string `json:"audience_rating"`
	AudienceScore  int    `json:"audience_score"`
}

type Posters struct {
	Thumbnail string `json:"thumbnail"`
	Profile   string `json:"Profile"`
	Detailed  string `json:"Detailed"`
	Original  string `json:"original"`
}

type Director struct {
	Name string `json:"name"`
}

type AlternateIds struct {
	IMDB string `json:"imdb"`
}

type Links struct {
	Self      string `json:"self"`
	Alternate string `json:"alternate"`
	Cast      string `json:"cast"`
	Clips     string `json:"clips"`
	Reviews   string `json:"reviews"`
	Similar   string `json:"similar"`
}

type Review struct {
	Critic      string      `json:"critic"`
	Date        string      `json:"date"`
	Freshness   string      `json:"freshness"`
	Publication string      `json:"publication"`
	Quote       string      `json:"quote"`
	Links       ReviewLinks `json:"links"`
}

type ReviewLinks struct {
	Review string `json:"review"`
}

type Clip struct {
	Title     string    `json:"title"`
	Duration  string    `json:"duration"`
	Thumbnail string    `json:"thumbnail"`
	Links     ClipLinks `json:"links"`
}

type ClipLinks struct {
	Alternate string `json:"alternate"`
}
