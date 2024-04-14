package main

import (
	"ass1/internal/data"
	"ass1/internal/handlers"
	"ass1/internal/service"
	"log"
	"net/http"
	"time"
)

func main() {
	m := data.New()
	s := service.New(m)
	h := handlers.New(s)
	srv := &http.Server{
		Addr:         ":8000",
		Handler:      h.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Println("server start")
	log.Fatal(srv.ListenAndServe())
}
