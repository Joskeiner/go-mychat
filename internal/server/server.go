package server

import (
	"log"
	"net/http"
	"time"
)

type Http struct {
	*http.Server
}

// this funtion running server
func (s *Http) Start() {
	log.Printf("Server is running on port %s", s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Panicf("Closed  server  error %s", err.Error())
	}
}

func NewServer(adr string, r http.Handler) *Http {
	return &Http{
		&http.Server{
			Addr:         adr,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      r,
		},
	}
}
