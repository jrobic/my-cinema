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
