package config

import (
	"github.com/jmoiron/sqlx"
)

// Global variables
var DB *sqlx.DB

// type Server struct {
// 	Cfg    Config
// 	Db     *sql.DB
// 	Router *chi.Mux
// }

func SetDbClient(db *sqlx.DB) {
	DB = db
}

// func (s *Server) GetRouter() *chi.Mux {
// 	return s.Router
// }
