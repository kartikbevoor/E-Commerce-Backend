package config

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	cfg    Config
	db     *sql.DB
	router *chi.Mux
}

func NewServer(cfg Config, db *sql.DB) *Server {
	return &Server{
		cfg:    cfg,
		db:     db,
		router: chi.NewRouter(),
	}
}

func (s *Server) Router() *chi.Mux {
	return s.router
}
