package service

import (
	"database/sql"
	"log"
	"time"

	"app/internal/model"
	"app/internal/repo"
	"app/pkg/errors"
)

type TaskService struct {
	Repo *repo.TaskRepo
}

func NewTaskService(r *repo.TaskRepo) *TaskService {
	return &TaskService{Repo: r}
}

func (s *TaskService) SubmitTask(req model.SubmitTaskReq) (*model.TaskRes, error) {
	if req.Name == "" {
		return nil, errors.NewUserError(4001, "task name can not be empty", nil)
	}

	task := model.Task{
		Name:      req.Name,
		DelayTime: req.DelayTime,
	}

	id, err := s.Repo.CreateTask(&task)

	if err != nil {
		return nil, errors.NewSystemError(5001, "Internal Server Error", err)
	}

	//单开协程完成主干任务

	go s.ProcessTaskAsync(id, req.DelayTime)

	return &model.TaskRes{
		TaskID: id,
		Status: "Submitted",
	}, nil
}

//单开协程的逻辑业务函数,这里只返回错误向前端传递错误，交给handler层，并且这一层和热破曾都不需要要http

func (s *TaskService) ProcessTaskAsync(id int64, delaytime int) {
	err := s.Repo.UpdateStatus(id, "running")
	if err != nil {
		log.Printf("[Service Error]failed to update running status %s", err)
		return
	}

	time.Sleep(time.Duration(delaytime) * time.Millisecond)

	err = s.Repo.UpdateStatus(id, "done")
	if err != nil {
		log.Printf("[Service Error]failed to update done status %s", err)
		return
	}
}

func (s *TaskService) GetTaskStatus(id int64) (*model.TaskRes, error) {
	if id <= 0 {
		return nil, errors.NewUserError(4002, "task id invalid", nil)
	}

	task, err := s.Repo.GetTaskStatus(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewUserError(4041, "task not found", err)
		}
		return nil, errors.NewSystemError(5002, "[Service Error]failed to get the taskid", err)
	}

	return &model.TaskRes{
		TaskID: task.ID,
		Status: task.Status,
		Result: task.Result,
	}, nil
}
