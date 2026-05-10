package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"myapp/internal/model"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisRepo结构体
type RedisRepo struct {
	redisClient *redis.Client
}

// 创建RedisRepo实例
func NewRedisRepo(client *redis.Client) *RedisRepo {
	return &RedisRepo{
		client,
	}
}

// 将任务加入队列
func (r *RedisRepo) Enqueue(ctx context.Context, id int64, name string, delaytime int) error {
	var msg *model.TaskRedis
	msg = &model.TaskRedis{
		Name:      name,
		ID:        id,
		DelayTime: delaytime,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return r.redisClient.LPush(ctx, "task:queue:normal", data).Err()
}

// 从队列中取出任务
func (r *RedisRepo) Dequeue(ctx context.Context) (*model.TaskRedis, error) {
	var msg model.TaskRedis

	res, err := r.redisClient.BRPop(ctx, 5*time.Second, "task:queue:normal").Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		return nil, err
	}

	err = json.Unmarshal([]byte(res[1]), &msg)
	return &msg, err
}

// 设置任务状态缓存
func (r *RedisRepo) SetStatusCache(ctx context.Context, id int64, status string) {
	key := fmt.Sprintf("key:task:%d", id)

	r.redisClient.Set(ctx, key, status, 5*time.Minute)
}

// 获取任务状态缓存
func (r *RedisRepo) GetStatusCache(ctx context.Context, id int64) (result string, err error) {
	key := fmt.Sprintf("key:task:%d", id)

	result, err = r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}

	return result, err

}
