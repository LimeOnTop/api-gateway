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

	app.Run(cfg, *devMode)
}