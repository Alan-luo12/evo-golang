package main

import (
	"os"
)

type Config struct {
	Port string
}

func Load_Config() *Config {
	port := os.Getenv("Port")
	if port == "" {
		port = "8080"
	}
	return &Config{
		Port: port,
	}
}
