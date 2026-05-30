package security

import (
	"context"
	"sync"
	"time"
)

// 抽象的Nonce存储接口
type NonceStore interface {
	UseNonceOnce(ctx context.Context, nonce string, ttl time.Duration) (bool, error)
}

// 本地内存存储实现NonceStore接口
type MemoryNonceStore struct {
	mu    sync.Mutex
	items map[string]time.Time
}

// 创建本地内存Nonce存储
func NewMemoryNonceStore() *MemoryNonceStore {
	return &MemoryNonceStore{
		items: make(map[string]time.Time),
	}
}

// 检查Nonce是否被使用过
// 如果未被使用过，返回true
// 如果已被使用过，返回false
// 如果ttl小于等于0，返回false
func (s *MemoryNonceStore) UseNonceOnce(ctx context.Context, nonce string, ttl time.Duration) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	for k, expireAt := range s.items {
		if now.After(expireAt) {
			delete(s.items, k)
		}
	}
	if expireAt, ok := s.items[nonce]; ok && now.Before(expireAt) {
		return false, nil
	}
	if ttl <= 0 {
		return false, nil
	}
	s.items[nonce] = now.Add(ttl)
	return true, nil
}
