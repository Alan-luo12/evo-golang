package config

import (
	"log"
	"os"
	"strconv"
)

//两个字段，端口号和数据库路径

type Config struct {
	Port      string
	DBPath    string
	RedisAddr string

	RateLimitCapacity   int64
	RateLimitRefillRate int64

	WorkerPool         int64
	JobQueue           int64
	ProcessConcurrency int64
}

// 加载配置文件
func Load_Config() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//获取数据库路径，如果环境变量不存在，则返回默认值
	dbpath := os.Getenv("DBPATH")
	if dbpath == "" {
		dbpath = "tasks.db"
	}

	//获取redis地址，如果环境变量不存在，则返回默认值
	redisaddr := os.Getenv("REDISADDR")
	if redisaddr == "" {
		redisaddr = "localhost:6379"
	}

	return &Config{
		Port:                port,
		DBPath:              dbpath,
		RedisAddr:           redisaddr,
		RateLimitCapacity:   Getenvint64("RATELIMITCAPACITY", 500),
		RateLimitRefillRate: Getenvint64("RATELIMITREFILLRATE", 100),
		WorkerPool:          Getenvint64("WORKERPOOLSIZE", 10),
		JobQueue:            Getenvint64("JOBQUEUESIZE", 100),
		ProcessConcurrency:  Getenvint64("PROCESSCONCURRENCY", 1),
	}
}

// 获取环境变量，如果环境变量不存在，则返回默认值
func Getenvint64(key string, devalue int64) int64 {
	v := os.Getenv(key)
	if v == "" {
		return devalue
	}

	result, err := strconv.ParseInt(v, 10, 64)
	if err != nil || result <= 0 {
		log.Printf("[warn] invalid value for %s env var, using default %d", key, devalue)
		return devalue
	}
	return result
}
