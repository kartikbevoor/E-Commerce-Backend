package server

import (
	"context"
	"ecommerce-backend/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	Cfg        config.Config
	Db         *sqlx.DB
	Router     *chi.Mux
	httpServer *http.Server
}

func main() {

	Server := NewServer()
	Server.Init()
	Server.Run()

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

	s.Db = config.NewDatabase(s.Cfg.Database)
	s.NewDatabase()
	config.SetDbClient(s.Db)
	config.SetConfig(s.Cfg)
}

func (s *Server) NewDatabase() {
	if s.Cfg.Database.Driver == "" {
		log.Fatal("please fill in database credentials in .env file or set in environment variable")
	}

	db, err := sqlx.Open(s.Cfg.Database.Driver, s.Cfg.Database.DSN)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatal(err)
	}

	s.Db = db

	// Set the below later
	// s.Db.SetMaxOpenConns(10)
	// s.Db.SetMaxIdleConns(10)
	// s.Db.SetConnMaxLifetime(10)
}

func (s *Server) Run() {
	s.httpServer = &http.Server{
		Addr:    ":" + s.Cfg.Server.Port,
		Handler: s.Router,
	}

	log.Printf("server starting on port %s", s.Cfg.Server.Port)

	go func() {
		start(s.httpServer)
	}()

	_ = gracefulShutdown(context.Background(), s)
}

func start(httpServer *http.Server) {

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}

func gracefulShutdown(ctx context.Context, s *Server) error {
	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-quit

	log.Println("Shutting down...")

	ctx, shutdown := context.WithTimeout(
		ctx,
		10*time.Second, // s.Cfg.Api.GracefulTimeout*time.Second,
	)
	defer shutdown()

	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}
	s.closeResources()

	return nil
}

func (s *Server) closeResources() { // ctx context.Context pass this as parameter for later resources
	if s.Db != nil {
		if err := s.Db.Close(); err != nil {
			log.Printf("error closing database: %v", err)
		}
	}
}
