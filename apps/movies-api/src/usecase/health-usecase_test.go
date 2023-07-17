package usecase_test

import (
	"testing"

	"github.com/jrobic/my-cinema/movies-api/src/domain"
	"github.com/jrobic/my-cinema/movies-api/src/usecase"
)

func TestHealth(t *testing.T) {
	t.Run("should return OK", func(t *testing.T) {
		want := domain.HealthResponse{
			Status: "OK",
		}

		usecases := usecase.NewAppUsecases()

		got := usecases.Health()

		if got.Status != want.Status {
			t.Errorf("got %q, want %q", got.Status, want.Status)
		}
	})
}
