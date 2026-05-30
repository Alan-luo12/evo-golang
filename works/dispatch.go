package works

import (
	"Lin/queue"
	"context"
)

// Dispatcher 任务分发器
// 从queue中dequeue job，然后将job发送到out channel
type Dispatcher struct {
	queue Queue
	out   chan queue.Job
}

// 这里给出定义实现queue和worker的依赖解耦，将queue和worker的实现分离，使它们可以独立开发和测试，
type Queue interface {
	Enqueue(ctx context.Context, job queue.Job) error
	Dequeue(ctx context.Context) (queue.Job, error)
}

// NewDispatcher 创建一个Dispatcher实例
func NewDispatcher(q Queue, buffer int) *Dispatcher {
	if buffer <= 0 {
		buffer = 100
	}
	return &Dispatcher{
		queue: q,
		out:   make(chan queue.Job, buffer),
	}
}

// Jobs 返回一个只读的jobchannel，用于接收dispatch的job，这是给workerpool使用的
func (d *Dispatcher) Jobs() <-chan queue.Job {
	return d.out
}

// Start 启动dispatch，从queue中dequeue job，然后将job发送到out channel
func (d *Dispatcher) Start(ctx context.Context) {
	go func() {
		//当出现ctx已取消或queue返回错误时，关闭out channel，目的是防止workerpool中的worker继续处理错误的任务，
		defer close(d.out)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				job, err := d.queue.Dequeue(ctx)
				//如果queue返回错误，继续下一个循环，目的是防止workerpool中的worker继续处理错误的任务，
				//在下一个循环中如果检测到ctx已取消，则直接返回，不继续处理错误的任务，
				if err != nil {
					continue
				}
				select {
				case d.out <- job:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
}
