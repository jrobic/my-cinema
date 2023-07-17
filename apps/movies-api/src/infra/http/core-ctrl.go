package httpserver

import "github.com/gofiber/fiber/v2"

type HealthResponse struct {
	Status string `json:"status"`
}

// @Summary Health check
// @Description Check if the API is up and running
// @Tags Health
// @Produce json
// @Router /health [get]
// @Success 200 {object} HealthResponse
func (s *MoviesAPIHttpServer) Health(c *fiber.Ctx) error {
	return c.JSON(s.usecases.Health())
}
