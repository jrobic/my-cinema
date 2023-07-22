package repository

import (
	"sync"

	"github.com/jrobic/my-cinema/movies-api/src/domain"
)

type MoviesInMemoryRepository struct {
	movies []*domain.Movie
	lock   sync.RWMutex
}

func NewMoviesInMemoryRepository(intialMovies []*domain.Movie) *MoviesInMemoryRepository {
	movies := []*domain.Movie{}

	movies = append(movies, intialMovies...)

	return &MoviesInMemoryRepository{
		movies: movies,
	}
}

func (r *MoviesInMemoryRepository) Insert(movie domain.Movie) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.movies = append(r.movies, &movie)

	return nil
}

func (r *MoviesInMemoryRepository) FindAll() ([]*domain.Movie, error) {
	return r.movies, nil
}

func (r *MoviesInMemoryRepository) InsertMany(movies []domain.Movie) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, movie := range movies {
		r.movies = append(r.movies, &movie)
	}

	return nil
}
