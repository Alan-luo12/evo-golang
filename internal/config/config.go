package config

import (
	"log"
	"os"
	"strconv"
)

//两个字段，端口号和数据库路径

type Config struct {
	Port       string
	SqlitePath string
	MysqlPath  string
	RedisAddr  string
	AuthToken  string
	MachineID  uint16

	LimitModel string

	RateLimitCapacity   int64
	RateLimitRefillRate int64

	DistLimitMax      int64
	DistLimitWindow   int64
	DistLimitFailOpen bool

	WorkerPool         int64
	JobQueue           int64
	ProcessConcurrency int64

	MaxConns        int64
	MaxIdleConns    int64
	ConnMaxLifetime int64
}

// 加载配置文件
func Load_Config() *Config {

	return &Config{
		Port:       Getenvstring("PORT", "8080"),
		SqlitePath: Getenvstring("DBPATH", "tasks.db"),
		MysqlPath:  Getenvstring("MYSQLPATH", "root:123456@tcp(localhost:3306)/data?charset=utf8mb4&parseTime=True&loc=Local"),
		RedisAddr:  Getenvstring("REDISADDR", "localhost:6379"),
		AuthToken:  Getenvstring("AUTHTOKEN", ""),
		MachineID:  uint16(Getenvint64("MACHINEID", 1)), //默认值为1，最大值为1023

		LimitModel: Getenvstring("MODEL", "dist"),

		DistLimitMax:    Getenvint64("DISTLIMITMAX", 200),
		DistLimitWindow: Getenvint64("DISTLIMITWINDOWMS", 1000),
		//是否开启当dist limit超过最大值时，是否返回错误
		DistLimitFailOpen: Getenvbool("DISTLIMITFAILOPEN", false),

		RateLimitCapacity:   Getenvint64("RATELIMITCAPACITY", 500),
		RateLimitRefillRate: Getenvint64("RATELIMITREFILLRATE", 100),

		WorkerPool:         Getenvint64("WORKERPOOLSIZE", 10),
		JobQueue:           Getenvint64("JOBQUEUESIZE", 100),
		ProcessConcurrency: Getenvint64("PROCESSCONCURRENCY", 5),

		MaxConns:        Getenvint64("MAXCONNS", 100),
		MaxIdleConns:    Getenvint64("MAXIDLECONNS", 100),
		ConnMaxLifetime: Getenvint64("CONNMAXLIFETIME", 5),
	}
}

// 获取环境变量，如果环境变量不存在，则返回默认值
func Getenvint64(key string, devalue int64) int64 {
	v := os.Getenv(key)
	if v == "" {
		log.Printf("[Info] %s env var not set, using default %d", key, devalue)
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

func Getenvbool(key string, devalue bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return devalue
	}

	switch v {
	case "true", "1", "yes", "y":
		return true

	case "false", "0", "no", "n":
		return false
	default:
		log.Printf("[Info] invalid value for %s env var, using default %v", key, devalue)
		return devalue
	}
}

func Getenvstring(key string, devalue string) string {
	v := os.Getenv(key)
	if v == "" {
		return devalue
	}
	return v
}
