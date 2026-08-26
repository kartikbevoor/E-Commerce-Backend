package config

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Cfg    Config
	Db     *sql.DB
	Router *chi.Mux
}

func NewServer(cfg Config, db *sql.DB) *Server {
	return &Server{
		Cfg:    cfg,
		Db:     db,
		Router: chi.NewRouter(),
	}
}

func (s *Server) GetRouter() *chi.Mux {
	return s.Router
}
