package service

import (
	"context"
	"log"
	"myapp/internal/model"
	"time"
)

// 启动worker池还有一个dispatchloop
func (s *TaskService) StartWorkers(ctx context.Context) {
	log.Printf("[Worker Started] the worker pool is starting with %d workers", s.workerpool)

	s.wg.Add(1)
	go s.dispatchloop(ctx)
	for i := int64(1); i <= s.workerpool; i++ {
		s.wg.Add(1)
		go s.Worker(ctx)
	}
}

func (s *TaskService) dispatchloop(ctx context.Context) {
	defer s.wg.Done()
	//diapatchloop是jobqueue的生产者负责从redis队列中取出任务，放入jobqueue中，负责关闭
	defer close(s.jobqueue)

	for {
		//避免在触发err的时候死循环，导致无法响应系统信号来优雅关闭，所以在每次循环开始的时候检查ctx是否已经被取消，如果已经被取消就退出循环
		if ctx.Err() != nil {
			log.Println("[Dispatch Loop] the dispatch loop is shutting down")
			return
		}

		msg, err := s.redisrepo.Dequeue(ctx)
		if err != nil {
			log.Println("[Dispatch Loop] error occurred while dequeuing task")
			time.Sleep(1 * time.Second)
			continue
		}

		if msg == nil {
			continue
		}

		select {
		//避免在没有msg的时候阻塞，导致无法响应系统信号来优雅关闭，所以使用select来监听ctx.Done()来判断是否需要退出
		case s.jobqueue <- msg:
		case <-ctx.Done():
			log.Println("[Dispatch Loop] the dispatch loop is shutting down")
			return
		}
	}

}

//消费队列开死循环，监听队列中是否有新的任务，如果有就取出任务，执行任务，执行完后更新任务状态，继续监听队列

func (s *TaskService) Worker(ctx context.Context) {
	defer s.wg.Done()
	log.Printf("[Workerd] started")
	defer log.Printf("[Worker] exited")

	//无需再监听ctx因为range在jobqueue被关闭的时候会自动退出循环，dispatchloop负责关闭jobqueue
	for msg := range s.jobqueue {
		func() {
			defer s.release()
			s.acquire()
			s.processtask(ctx, msg.ID, msg.Name, msg.DelayTime)
		}()
	}

}

//按序排列方便阅读

func (s *TaskService) acquire() {
	s.processconcurrency <- struct{}{}
}

func (s *TaskService) processtask(ctx context.Context, id int64, name string, delaytime int) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()
	t := model.Task{
		ID:        id,
		Name:      name,
		Status:    "running",
		DelayTime: delaytime,
	}

	_, err := s.repo.CreateTask(ctx, id, &t)
	if err != nil {
		log.Printf("[Error] failed to create task in db %s", err)
		return
	}

	s.redisrepo.SetStatusCache(ctx, id, "running")

	//关键逻辑，模拟任务执行的时间，实际中这里可能是调用第三方接口，或者执行一些复杂的计算等
	time.Sleep(time.Duration(delaytime) * time.Millisecond)

	err = s.repo.UpdateStatus(ctx, id, "done")
	if err != nil {
		log.Printf("[Error] failed to update task status in db %s", err)
		err = s.repo.UpdateStatus(ctx, id, "failed")
		s.redisrepo.SetStatusCache(ctx, id, "failed")
		if err != nil {
			log.Printf("[Error] failed to update task status to failed in db %s", err)
			return
		}
		return
	}

	s.redisrepo.SetStatusCache(ctx, id, "done")
}

func (s *TaskService) release() {
	<-s.processconcurrency
}

// 等待所有的worker协程退出，通常在main函数中调用这个函数来等待所有的worker协程退出，确保优雅关闭
func (s *TaskService) Wait() {
	s.wg.Wait()
	log.Printf("[Worker Wait] all workers have exited")
}
