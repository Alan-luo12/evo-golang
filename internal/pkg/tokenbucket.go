package pkg

import (
	"sync"
	"time"
)

// 令牌桶结构体
type TokenBucket struct {
	capacity   int64
	tokens     int64
	refillrate int64
	lastrefill time.Time
	mu         sync.Mutex
}

// 创建令牌桶
func NewTokenBucket(capacity int64, refillrate int64) *TokenBucket {

	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillrate: refillrate,
		lastrefill: time.Now(),
	}
}

// 检查是否有足够的令牌，如果有则消耗一个令牌并返回true，否则返回false
func (tb *TokenBucket) AllowLocal() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.lastrefill).Nanoseconds()

	newtokens := elapsed * tb.refillrate / int64(time.Second)

	//根据时间间隔和补充速率计算当前令牌数量，统一int64类型
	tb.tokens += newtokens

	//如果令牌数量超过容量，重置为容量
	if tb.tokens <= 0 {
		return false
	}

	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	//更新上次补充时间
	tb.lastrefill = tb.lastrefill.Add(time.Duration(newtokens * int64(time.Second) / tb.refillrate))

	//消耗一个令牌
	tb.tokens--

	return true
}
