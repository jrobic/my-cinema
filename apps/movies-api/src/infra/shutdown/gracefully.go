package shutdown

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Gracefully() {
	quit := make(chan os.Signal, 1)
	defer close(quit)

	log.Printf("gracefully shutdown started")

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	s := <-quit

	log.Printf("received signal %s", s)
}
