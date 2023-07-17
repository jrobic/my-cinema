package main

import (
	"log"
	"os"

	"github.com/jrobic/my-cinema/movies-api/src/config"
	httpserver "github.com/jrobic/my-cinema/movies-api/src/infra/http"
	"github.com/jrobic/my-cinema/movies-api/src/infra/shutdown"
)

// @title Movies API
// @description This is a movies API.
// @version 0.1

// @contact.name Jonathan Robic
// @contact.email hello@jonathanrobic.fr

// @host localhost:3001
// @BasePath /
// @schemes http
func main() {
	// setup exit code for graceful shutdown
	var exitCode int

	defer func() {
		log.Printf("exiting with code %d", exitCode)
		os.Exit(exitCode)
	}()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("could not load config: %v", err)
		exitCode = 1
		return
	}

	server, err := httpserver.NewMoviesAPIHttpServer()
	defer server.Cleanup()

	if err != nil {
		log.Printf("could not create server: %v", err)
		exitCode = 1
		return
	}

	go func() {
		if err = server.App.Listen(":" + cfg.Port); err != nil {
			log.Fatalf("could not listen on port %s %v", cfg.Port, err)
		}
	}()

	shutdown.Gracefully()

}
