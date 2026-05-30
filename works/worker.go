package works

import (
	"Lin/queue"
	"context"
	"log"
	"sync"
)

type Pool struct {
	size      int
	jobs      <-chan queue.Job
	processor Processor
	wg        sync.WaitGroup
}

func NewWorkerPool(size int, jobs <-chan queue.Job, processor Processor) *Pool {
	if size <= 0 {
		size = 1
	}
	return &Pool{
		size:      size,
		jobs:      jobs,
		processor: processor,
	}
}

// start 启动worker池
// 如果ctx已取消，返回错误
func (p *Pool) Start(ctx context.Context) {
	//workerID从0开始
	workerID := 0
	//启动worker池
	for i := 0; i < p.size; i++ {
		workerID++
		p.wg.Add(1)
		//启动一个worker，从job通道中接收任务，调用处理器处理任务，处理完成后，将结果写入结果通道
		go func(workerID int) {
			defer p.wg.Done()
			for {
				//这里提供了详细的错误处理，包括ctx已取消，job通道已关闭，处理器返回错误
				select {
				case <-ctx.Done():
					log.Printf("[Worker] worker=%d stopped", workerID)
					return
				case job, ok := <-p.jobs:
					if !ok {
						log.Printf("[Worker] worker=%d job channel closed", workerID)
						return
					}
					if err := p.processor.Process(ctx, job); err != nil {
						log.Printf("[Worker] worker=%d job=%s error=%v", workerID, job.ID, err)
					}
				}
			}
		}(workerID)
	}
}

// 等待所有worker完成
func (p *Pool) Wait() {
	p.wg.Wait()
}
