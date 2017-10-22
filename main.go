package main

import (
	"fmt"
	"github.com/octoberstorm/angular2-kanban/server/config"
	"github.com/octoberstorm/angular2-kanban/server/controllers"
	"log"
	"net/http"
	"time"
)

func main() {
	c := config.Get()
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	s := &http.Server{
		Addr:         addr,
		Handler:      controllers.Router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Printf("Listening on %s", addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
