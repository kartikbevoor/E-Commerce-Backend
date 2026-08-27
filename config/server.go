package config

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

// Global variables
var DB *sqlx.DB

type Server struct {
	Cfg    Config
	Db     *sql.DB
	Router *chi.Mux
}

func SetDbClient(db *sqlx.DB) {
	DB = db
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
