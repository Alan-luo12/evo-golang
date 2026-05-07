package setup

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func InitResdis(addr string) *redis.Client {
	Client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := Client.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("[Error] Failed to connect to Redis at %s", addr)
	}

	log.Printf("[Init] Succedd to connect to Redis at %s", addr)

	return Client
}
