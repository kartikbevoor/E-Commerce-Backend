package server

import (
	"ecommerce-backend/config"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	cfg    config.Config
	Router chi.Router
}

func main() {
	//Server := NewServer()
}

func NewServer() *Server {
	return &Server{
		cfg:    *config.GetConfigurations(),
		Router: chi.NewRouter(),
	}
}
