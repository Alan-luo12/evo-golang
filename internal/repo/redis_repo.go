package repo

import (
	"app/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	RedisClient *redis.Client
}

func NewRedisRepo(client *redis.Client) *RedisRepo {
	return &RedisRepo{
		client,
	}
}

func (r *RedisRepo) Enqueue(ctx context.Context, taskid int64, delaytime int) error {
	var msg *model.TaskRedis
	msg = &model.TaskRedis{
		ID:        taskid,
		DelayTime: delaytime,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return r.RedisClient.LPush(ctx, "task:queue:normal", data).Err()
}

func (r *RedisRepo) Dequeue(ctx context.Context) (*model.TaskRedis, error) {
	var msg model.TaskRedis

	res, err := r.RedisClient.BRPop(ctx, 5*time.Second, "task:queue:normal").Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		return nil, err
	}

	err = json.Unmarshal([]byte(res[1]), &msg)
	return &msg, err
}

func (r *RedisRepo) SetStatusCache(ctx context.Context, id int64, status string) {
	key := fmt.Sprintf("key:task:%d", id)

	r.RedisClient.Set(ctx, key, status, 5*time.Minute)
}

func (r *RedisRepo) GetStatusCache(ctx context.Context, id int64) (result string, err error) {
	key := fmt.Sprintf("key:task:%d", id)

	result, err = r.RedisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}

	return result, err

}
