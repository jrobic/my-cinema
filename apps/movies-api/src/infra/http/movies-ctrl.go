package httpserver

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (s *MoviesAPIHttpServer) IngestMoviesFile(c *fiber.Ctx) error {
	data, err := s.usecases.IngestMoviesFile("./data/movies.csv")
	// data, err := s.usecases.IngestMoviesFile("./src/usecase/__tests__/movies.csv")

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": data,
	})

}

func (s *MoviesAPIHttpServer) FetchAll(c *fiber.Ctx) error {
	movies, err := s.usecases.FindAll()

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": movies,
	})
}
