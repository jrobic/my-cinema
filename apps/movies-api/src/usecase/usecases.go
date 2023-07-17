package usecase

import "github.com/jrobic/my-cinema/movies-api/src/domain"

type Usecases interface {
	Health() *domain.HealthResponse
}

type AppUsecases struct{}

func NewAppUsecases() *AppUsecases {
	return &AppUsecases{}
}
