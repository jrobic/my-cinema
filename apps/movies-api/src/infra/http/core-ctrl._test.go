package httpserver_test

import (
	"testing"

	motesting "github.com/jrobic/my-cinema/movies-api/src"
	"github.com/jrobic/my-cinema/movies-api/src/domain"
	httpserver "github.com/jrobic/my-cinema/movies-api/src/infra/http"
)

func TestHealthCtrl(t *testing.T) {
	t.Run("should return OK", func(t *testing.T) {
		want := domain.HealthResponse{
			Status: "OK",
		}

		server, _ := httpserver.NewMoviesAPIHttpServer()

		req := motesting.NewRequest("GET", "/health", nil)

		response, _ := server.App.Test(req, -1)

		got := motesting.ParseJSONReponse(t, response.Body, domain.HealthResponse{})

		if got.Status != want.Status {
			t.Errorf("got %q, want %q", got.Status, want.Status)
		}
	})
}
