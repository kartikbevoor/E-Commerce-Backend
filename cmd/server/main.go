package server

import (
	"ecommerce-backend/config"
	"log"
	"net/http"
)

func main() {
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
		Handler: app.Router(),
	}

	log.Printf("server starting on port %s", cfg.Server.Port)

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
