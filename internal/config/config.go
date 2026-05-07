package config

import (
	"os"
)

//两个字段，端口号和数据库路径

type Config struct {
	Port      string
	DBPath    string
	RedisAddr string
}

func Load_Config() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	DBpath := os.Getenv("DBPATH")
	if DBpath == "" {
		DBpath = "tasks.db"
	}

	redisaddr := os.Getenv("RedisAddr")
	if redisaddr == "" {
		redisaddr = "localhost:6379"
	}

	return &Config{
		Port:      port,
		DBPath:    DBpath,
		RedisAddr: redisaddr,
	}
}
