package main

import (
	"flag"
	"log"

	"api-gateway/config"
	"api-gateway/internal/app"
)

func main() {
	devMode := flag.Bool("dev", false, "Run in development mode")
	flag.Parse()
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	srv, err := app.NewServer(cfg)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	if err := srv.Run(cfg, *devMode); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}