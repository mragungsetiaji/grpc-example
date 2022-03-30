package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	s "github.com/mragungsetiaji/grpc-example/server"
)

func init() {
	s.LoadConfig()
}

// Entry point of the scheduler application.
func main() {

	go s.API()
	go s.StartGRPCServer()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case s := <-sig:
			log.Fatalf("Signal (%d) received, stopping\n", s)
		}
	}
}
