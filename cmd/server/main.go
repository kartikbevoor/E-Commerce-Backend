package server

import (
	"database/sql"
	"ecommerce-backend/config"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Cfg    config.Config
	Db     *sql.DB
	Router *chi.Mux
}

func main() {

	//Server := NewServer()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	db, err := config.NewDatabase(cfg.Database)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.Close()

	app := config.NewServer(cfg, db)

	httpServer := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: app.GetRouter(),
	}

	log.Printf("server starting on port %s", cfg.Server.Port)

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}

func NewServer() *Server {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	return &Server{
		Cfg:    cfg,
		Router: chi.NewRouter(),
	}
}

func (s *Server) Init() {
	s.NewDatabase()
}

func (s *Server) NewDatabase() {
	if s.Cfg.Database.Driver == "" {
		log.Fatal("please fill in database credentials in .env file or set in environment variable")
	}

	db, err := sql.Open(s.Cfg.Database.Driver, s.Cfg.Database.DSN)
	if err != nil {
		log.Fatal(err)
	}

	s.Db = db

	// Set the below later
	// s.Db.SetMaxOpenConns(10)
	// s.Db.SetMaxIdleConns(10)
	// s.Db.SetConnMaxLifetime(10)
}
