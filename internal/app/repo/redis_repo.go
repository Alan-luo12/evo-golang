package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"myapp/internal/model"
	"strconv"
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
func (r *RedisRepo) SetStatusCache(ctx context.Context, id int64, status string) error {
	key := fmt.Sprintf("cache:task:%d", id)

	err := r.redisClient.Set(ctx, key, status, 5*time.Minute).Err()
	return err
}

// 获取任务状态缓存
func (r *RedisRepo) GetStatusCache(ctx context.Context, id int64) (result string, err error) {
	key := fmt.Sprintf("cache:task:%d", id)

	result, err = r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}

	return result, err

}

// 11
// 11
func (r *RedisRepo) UseNonceOnce(ctx context.Context, nonce string, ttl time.Duration) (bool, error) {
	key := fmt.Sprintf("security:nounce:%s", nonce)
	return r.redisClient.SetNX(ctx, key, "1", ttl).Result()
}

var distlimitscript = redis.NewScript(`
	local current = redis.call("INCR",KEYS[1])
	if current == 1 then
		redis.call("PEXPIRE",KEYS[1],ARGV[1])
	end

	local ttl = redis.call("PTTL",KEYS[1])
	
	if current > tonumber(ARGV[2]) then
		return {0, current, ttl}
	end
	return {1, current, ttl}
`)

func (r *RedisRepo) AllowDist(ctx context.Context, key string, expire time.Duration, max int64) (*model.DistRes, error) {
	raw, err := distlimitscript.Run(ctx,
		r.redisClient,
		[]string{key},
		expire.Milliseconds(), max,
	).Result()

	if err != nil {
		return nil, err
	}

	res, ok := raw.([]interface{})
	if !ok || len(res) != 3 {
		return nil, err
	}

	allow, ok1 := ToInt64(res[0])
	current, ok2 := ToInt64(res[1])
	ttl, ok3 := ToInt64(res[2])

	if !ok1 || !ok2 || !ok3 {
		return nil, err
	}

	return &model.DistRes{
		Allow:   allow == 1,
		Current: current,
		TTL:     time.Duration(ttl) * time.Millisecond,
	}, nil
}

func ToInt64(v interface{}) (int64, bool) {
	switch t := v.(type) {
	case int64:
		return t, true
	case int:
		return int64(t), true
	case string:
		int64, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			return 0, false
		}
		return int64, true
	default:
		return 0, false
	}

}
