package setup

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

//初始化Redis客户端
func InitResdis(addr string) *redis.Client {
	//创建Redis客户端
	Client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	//ping一下，检查是否连接成功
	if err := Client.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("[Error] Failed to connect to Redis at %s", addr)
	}

	log.Printf("[Init] Succedd to connect to Redis at %s", addr)

	return Client
}
