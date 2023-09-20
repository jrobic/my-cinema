package domain

import (
	"strconv"
	"strings"

	"github.com/gofrs/uuid/v5"
)

type Movie struct {
	ID                  uuid.UUID `json:"id" bson:"_id"`
	TmdbID              string    `json:"tmdbId" bson:"tmdbId"`
	Title               string    `json:"title" bson:"title"`
	Overview            string    `json:"overview" bson:"overview"`
	OriginalLanguage    string    `json:"originalLanguage" bson:"originalLanguage"`
	Genres              []string  `json:"genres" bson:"genres"`
	ReleaseDate         string    `json:"releaseDate" bson:"releaseDate"`
	Popularity          float64   `json:"popularity" bson:"popularity"`
	ProductionCompanies []string  `json:"productionCompanies" bson:"productionCompanies"`
	Budget              int64     `json:"budget" bson:"budget"`   // revenue in centimes
	Revenue             int64     `json:"revenue" bson:"revenue"` // revenue in centimes
	Runtime             int64     `json:"runtime" bson:"runtime"`
	Status              string    `json:"status" bson:"status"`
	Tagline             string    `json:"tagline" bson:"tagline"`
	VoteAverage         float64   `json:"voteAverage" bson:"voteAverage"`
	VoteCount           int64     `json:"voteCount" bson:"voteCount"`
	Credits             []string  `json:"credits" bson:"credits"`
	Keywords            []string  `json:"keywords" bson:"keywords"`
	PosterPath          string    `json:"posterPath" bson:"posterPath"`
	BackdropPath        string    `json:"backdropPath" bson:"backdropPath"`
}

// idRecord                  = iota
// titleRecord               // 1
// genresRecord              // 2
// originalLanguageRecord    // 3
// overviewRecord            // 4
// popularityRecord          // 5
// productionCompaniesRecord // 6
// releaseDateRecord         // 7
// budgetRecord              // 8
// revenueRecord             // 9
// runtimeRecord             // 10
// statusRecord              // 11
// taglineRecord             // 12
// voteAverageRecord         // 13
// voteCountRecord           // 14
// creditsRecord             // 15
// keywordsRecord            // 16
// posterPathRecord          // 17
// backdropPathRecord        // 18

func NewMovieFromCSVRecord(rec []string) Movie {

	popularity, _ := strconv.ParseFloat(rec[5], 64)
	budget, _ := strconv.ParseFloat(rec[8], 64)
	revenue, _ := strconv.ParseFloat(rec[9], 64)
	runtime, _ := strconv.ParseFloat(rec[10], 64)
	voteAverage, _ := strconv.ParseFloat(rec[13], 32)
	voteCount, _ := strconv.ParseFloat(rec[14], 64)

	return Movie{
		ID:                  uuid.Must(uuid.NewV4()),
		TmdbID:              rec[0],
		Title:               rec[1],
		Overview:            rec[4],
		OriginalLanguage:    rec[3],
		Genres:              strings.Split(rec[2], "-"),
		ReleaseDate:         rec[7],
		Popularity:          popularity,
		ProductionCompanies: strings.Split(rec[6], "-"),
		Budget:              int64(budget * 100.0),
		Revenue:             int64(revenue * 100.0),
		Runtime:             int64(runtime),
		Status:              rec[11],
		Tagline:             rec[12],
		VoteAverage:         float64(int(voteAverage*10)) / 10,
		VoteCount:           int64(voteCount),
		Credits:             strings.Split(rec[15], "-"),
		Keywords:            strings.Split(rec[16], "-"),
		PosterPath:          rec[17],
		BackdropPath:        rec[18],
	}
}

type MovieRepository interface {
	Insert(movie Movie) error
	FindAll() ([]*Movie, error)
	InsertMany(movies []Movie) error
}
