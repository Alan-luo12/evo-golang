package queue

import "context"

type Job struct {
	ID      string
	Name    string
	Payload []byte
}
type Queue interface {
	Enqueue(ctx context.Context, job Job) error
	Dequeue(ctx context.Context) (Job, error)
}
