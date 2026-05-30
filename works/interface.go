package works

import (
	"Lin/queue"
	"context"
)

/*
   依赖关系
   ┌─────────────────┐
   │   MemoryQueue   │
   │  (内置实现)      │
   └────────┬────────┘
            │ 实现
   ┌────────▼────────┐
   │   Queue 接口     │
   │  (queue包定义)   │
   └────────┬────────┘
            │ 依赖
   ┌────────▼────────┐
   │   Dispatcher    │
   │  (需要Queue)     │
   └────────┬────────┘
            │ 给出
   ┌────────▼────────┐
   │  Jobs() 方法     │
   │  返回只读通道     │
   └────────┬────────┘
            │ 依赖
   ┌────────▼────────┐
   │      worker     │
   │  (需要jobs通道)  │
   └────────┬────────┘
            │
   ┌────────▼────────┐
   │    Processor    │
   │  (业务实现)      │
   └─────────────────┘
*/

// Processor 任务处理器接口
type Processor interface {
	Process(ctx context.Context, job queue.Job) error
}

// ProcessFunc 提供了一个简便的实现，直接返回函数调用的结果
type ProcessFunc func(ctx context.Context, job queue.Job) error

// 类似于适配器模式，将函数转换为Processor接口的实现，和go中的handler和handelrfunc思想一致
func (p ProcessFunc) Process(ctx context.Context, job queue.Job) error {
	return p(ctx, job)
}
