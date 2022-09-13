package main

import (
	"context"
	"github.com/alfiehiscox/go-practice-service/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, " go-practice-service", log.LstdFlags)

	// Handler
	hh := handlers.NewHello(l)

	// ServeMux on index of port
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	// Server Configuration
	s := http.Server{
		Addr:         ":8080",
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Using goroutine to handle exiting gracefully
	go func() {
		l.Println("Started listening on Port: 8080")
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Server Info: %s\n", err)
			os.Exit(1)
		}
	}()

	// Make Channel to listen to signals on the goroutine
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Listen to the channel and exit gracefully
	sig := <-c
	l.Println("Got Signal: ", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)

}
