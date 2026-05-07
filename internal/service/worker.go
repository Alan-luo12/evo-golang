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
		Status:    "pending",
		DelayTime: delaytime,
	}
	s.repo.CreateTask(&t)

	s.repo.UpdateStatus(id, "running")
	s.queuerepo.SetStatusCache(ctx, id, "running")

	time.Sleep(time.Duration(delaytime) * time.Millisecond)

	s.repo.UpdateStatus(id, "done")
	s.queuerepo.SetStatusCache(ctx, id, "done")
}
