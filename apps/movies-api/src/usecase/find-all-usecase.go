package usecase

import "github.com/jrobic/my-cinema/movies-api/src/domain"

func (au *AppUsecases) FindAll() ([]*domain.Movie, error) {
	return au.MoviesRepo.FindAll()
}
