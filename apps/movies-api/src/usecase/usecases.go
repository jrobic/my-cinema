package usecase

import "github.com/jrobic/my-cinema/movies-api/src/domain"

type Usecases interface {
	Health() *domain.HealthResponse
	IngestMoviesFile(file string) (IngestedMoviesFile, error)
	FindAll() ([]*domain.Movie, error)
}

type AppUsecases struct {
	MoviesRepo domain.MovieRepository
}

func NewAppUsecases(moviesRepo domain.MovieRepository) *AppUsecases {
	return &AppUsecases{
		MoviesRepo: moviesRepo,
	}
}
