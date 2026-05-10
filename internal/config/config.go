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
	AuthToken string

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

	authtoken := os.Getenv("AUTHTOKEN")
	if authtoken == "" {
		authtoken = ""
	}

	return &Config{
		Port:                port,
		DBPath:              dbpath,
		RedisAddr:           redisaddr,
		AuthToken:           authtoken,
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
	//将字符串转换为int64
	result, err := strconv.ParseInt(v, 10, 64)
	//转换失败或者结果小于零就采用默认值
	if err != nil || result <= 0 {
		log.Fatalf("[Fatal] invalid value for %s env var, using default %d", key, devalue)
	}
	return result
}
