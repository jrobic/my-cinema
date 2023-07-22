package usecase

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/jrobic/my-cinema/movies-api/src/domain"
)

const (
	id = iota
	title
	genres
	originalLanguage
	overview
	popularity
	productionCompanies
	releaseDate
	budget
	revenue
	runtime
	status
	tagline
	voteAverage
	voteCount
	credits
	keywords
	posterPath
	backdropPath
)

type IngestedMoviesFile struct {
	Count int
}

func (au *AppUsecases) IngestMoviesFile(file string) (IngestedMoviesFile, error) {
	f, err := os.Open(file)
	if err != nil {
		return IngestedMoviesFile{}, err
	}

	csvReader := csv.NewReader(f)
	// Read first line to skip header
	csvReader.Read()

	movies := []domain.Movie{}

	for {
		rec, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return IngestedMoviesFile{}, err
		}

		movies = append(movies, domain.NewMovieFromCSVRecord(rec))
	}

	f.Close()

	err = au.MoviesRepo.InsertMany(movies)
	if err != nil {
		return IngestedMoviesFile{}, err
	}

	return IngestedMoviesFile{Count: len(movies)}, nil
}

// Create as many "shufflings" as you can!
// With input 'abc':
// Your function should return ['abc','acb','bac','bca','cab','cba']
func Permutations(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}

	perms := []string{}
	for i, c := range s {
		// Remove the character at index i from the string
		// and calculate all the permutations of the remaining characters
		// (i.e. recursive call).
		// Then, add the removed character to the front of each permutation.
		// This will give us all the permutations that start with that character.
		// Finally, append these permutations to our running list of permutations.
		perms = append(perms, prependToAll(string(c), Permutations(s[:i]+s[i+1:]))...)
	}

	return perms
}

func prependToAll(c string, ss []string) []string {
	for i, s := range ss {
		ss[i] = c + s
	}

	return ss
}
