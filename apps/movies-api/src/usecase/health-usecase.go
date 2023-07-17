package usecase

import "github.com/jrobic/my-cinema/movies-api/src/domain"

func (au *AppUsecases) Health() *domain.HealthResponse {
	return &domain.HealthResponse{
		Status: "OK",
	}
}
