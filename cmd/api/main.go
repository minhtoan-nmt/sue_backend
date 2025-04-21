package main

import (
	"context"
	"fmt"
	"log"
	"sue_backend/config"
	"sue_backend/internal/infra/db"
	"sue_backend/internal/transport/http"
	"sue_backend/pkg/logger"
)

func main() {

	ctx := context.Background()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cfg, err := config.LoadAllConfigs("config")
	log := logger.New(cfg.Log)    // zap/sl
	router := http.NewRouter(cfg) // wires handlers + middleware

	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	postgresDB, err := db.InitDB(ctx, cfg.DB)
	if err != nil {
		log.Fatalf("failed to connect to Postgres: %v", err)
	}
	defer postgresDB.Close()

	PGStore := db.NewPostgresStore(postgresDB)

	rows, err := PGStore.ExecQuery(ctx, "SELECT * FROM student LIMIT 5")
	if err != nil {
		log.Fatal("Query failed:", err)
	}

	for _, row := range rows {
		fmt.Println(row)
	}

	log.Infof("⇢ starting HTTP server on %s", cfg.App.HostPort)
	if err := router.Run(cfg.App.HostPort); err != nil {
		log.Fatalf("server error: %v", err)
	}

}
