package queue

import (
	"context"
)

// MemoryQueue 内存队列
type MemoryQueue struct {
	ch chan Job
}

// NewMemoryQueue 创建内存队列
func NewMemoryQueue(size int) *MemoryQueue {
	if size <= 0 {
		size = 100
	}
	return &MemoryQueue{
		ch: make(chan Job, size),
	}
}

// Enqueue 入队
// 如果ctx已取消，返回错误
func (q *MemoryQueue) Enqueue(ctx context.Context, job Job) error {
	select {
	case q.ch <- job:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Dequeue 出队
// 如果ctx已取消，返回错误
func (q *MemoryQueue) Dequeue(ctx context.Context) (Job, error) {
	select {
	case job := <-q.ch:
		return job, nil
	case <-ctx.Done():
		return Job{}, ctx.Err()
	}
}
