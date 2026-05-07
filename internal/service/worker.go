package service

import (
	"app/internal/model"
	"context"
	"log"
	"time"
)

func (s *TaskService) Worker(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			log.Printf("[Shutting Dowwn]the Worker is shutting down")
			return

		default:

			msg, err := s.queuerepo.Dequeue(ctx)
			if err != nil {
				log.Printf("[Error] find error in redis %s", err)
				time.Sleep(1 * time.Second)
			}

			if msg == nil {
				continue
			}

			s.processtask(ctx, msg.ID, msg.Name, msg.DelayTime)
		}
	}
}

func (s *TaskService) processtask(ctx context.Context, id int64, name string, delaytime int) {

	t := model.Task{
		ID:        id,
		Name:      name,
		Status:    "running",
		DelayTime: delaytime,
	}
	_, err := s.repo.CreateTask(&t)
	if err != nil {
		log.Printf("[Error] failed to create task in db %s", err)
		return
	}

	s.queuerepo.SetStatusCache(ctx, id, "running")

	//关键逻辑，模拟任务执行的时间，实际中这里可能是调用第三方接口，或者执行一些复杂的计算等
	time.Sleep(time.Duration(delaytime) * time.Millisecond)

	err = s.repo.UpdateStatus(id, "done")
	if err != nil {
		log.Printf("[Error] failed to update task status in db %s", err)
		err = s.repo.UpdateStatus(id, "failed")
		s.queuerepo.SetStatusCache(ctx, id, "failed")
		if err != nil {
			log.Printf("[Error] failed to update task status to failed in db %s", err)
			return
		}
		return
	}

	s.queuerepo.SetStatusCache(ctx, id, "done")
}
