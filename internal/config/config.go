package config

import (
	"database/sql"
	"flag"
	"log"

	"github.com/wellls/devx/pkg/database"
)

type Config struct {
	ServerAddress int
	DB            *sql.DB
}

func LoadConfig() *Config {
	var cfg Config

	flag.IntVar(&cfg.ServerAddress, "port", 3333, "Server port to listen on")
	flag.Parse()

	// start database connection
	db, err := database.NewConnection()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	cfg.DB = db

	return &cfg
}
