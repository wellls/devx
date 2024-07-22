package config

import (
	"flag"
)

type Config struct {
	ServerAddress int
}

func LoadConfig() *Config {
	var cfg Config

	flag.IntVar(&cfg.ServerAddress, "port", 3333, "Server port to listen on")
	flag.Parse()

	return &cfg
}
