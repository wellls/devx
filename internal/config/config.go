package config

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress int
}

func LoadConfig() *Config {
	var cfg Config

	flag.IntVar(&cfg.ServerAddress, "port", 3333, "Server port to listen on")
	flag.Parse()

	// Load env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file", err)
	}

	return &cfg
}
