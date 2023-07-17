package httpserver

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Timer will measure how long it takes before a response is returned
func Timer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// start timer
		start := time.Now()
		// next routes
		err := c.Next()
		// stop timer
		stop := time.Now()
		// Do something with response
		c.Append("Server-Timing", fmt.Sprintf("%v", stop.Sub(start).String()))
		// return stack error if exist
		return err
	}
}
