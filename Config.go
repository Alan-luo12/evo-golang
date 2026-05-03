package main

import (
	"os"
)

type Config struct {
	Port   string
	DBPath string
}

func Load_Config() *Config {
	port := os.Getenv("Port")
	dbpath := os.Getenv("DBPath")

	if port == "" {
		port = "8080"
	}
	if dbpath == "" {
		dbpath = "tasks.db"
	}

	return &Config{
		Port:   port,
		DBPath: dbpath,
	}
}
