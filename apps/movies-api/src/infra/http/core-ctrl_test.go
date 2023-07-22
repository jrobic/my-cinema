package httpserver_test

import (
	"testing"

	motesting "github.com/jrobic/my-cinema/movies-api/src"
	"github.com/jrobic/my-cinema/movies-api/src/domain"
	httpserver "github.com/jrobic/my-cinema/movies-api/src/infra/http"
	"github.com/jrobic/my-cinema/movies-api/src/infra/repository"
)

func TestHealthCtrl(t *testing.T) {
	t.Run("should return OK", func(t *testing.T) {
		want := domain.HealthResponse{
			Status: "OK",
		}

		moviesRepo := repository.NewMoviesInMemoryRepository([]*domain.Movie{})
		serverDeps := httpserver.MoviesAPIHttpServerDeps{
			MoviesRepo: moviesRepo,
		}
		server, _ := httpserver.NewMoviesAPIHttpServer(serverDeps)

		req := motesting.NewRequest("GET", "/health", nil)

		response, _ := server.App.Test(req, -1)

		got := motesting.ParseJSONReponse(t, response.Body, domain.HealthResponse{})

		if got.Status != want.Status {
			t.Errorf("got %q, want %q", got.Status, want.Status)
		}
	})
}
