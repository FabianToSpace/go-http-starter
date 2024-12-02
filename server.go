package main

import (
	"log"
	"net/http"
)

type Server struct {
	serveMux http.ServeMux
	logf     func(f string, v ...interface{})
	db       *DatabaseConnection
}

func newServer() (*Server, error) {
	server := &Server{
		logf: log.Printf,
	}

	db, err := NewDatabaseConnection()
	if err != nil {
		server.logf("error in setting up database connection: %v", err)
		return nil, err
	}
	server.db = db

	server.registerRoutes()

	return server, nil
}

func (s *Server) registerRoutes() {
	s.serveMux.HandleFunc("/api/health", s.healthHandler)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.serveMux.ServeHTTP(w, r)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Service is alive"))
}
