package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	w "github.com/mragungsetiaji/grpc-example/worker"
)

func init() {
	w.LoadConfig()
}

// Entry point of the worker application.
func main() {

	go w.StartGRPCServer()
	go w.RegisterWorker()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case s := <-sig:
			fatal(fmt.Sprintf("Signal (%d) received, stopping\n", s))
		}
	}
}

func fatal(message string) {
	w.DeregisterWorker()
	log.Fatalln(message)
}
