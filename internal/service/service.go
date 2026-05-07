package service

import (
	"context"
	"database/sql"

	"app/internal/model"
	"app/internal/repo"
	"app/pkg/errors"
	"app/pkg/snowid"
)

type TaskService struct {
	repo      *repo.TaskRepo
	queuerepo *repo.RedisRepo
}

func NewTaskService(ra *repo.TaskRepo, rb *repo.RedisRepo) *TaskService {
	return &TaskService{
		repo:      ra,
		queuerepo: rb,
	}
}

func (s *TaskService) SubmitTask(req model.TaskSubmit) (*model.TaskRes, error) {

	i, err := snowid.NextID()
	id := int64(i)
	if err != nil {
		return nil, errors.NewSystemError(5003, "snowid error", err)
	}

	//这里接受的是一个submittaskreq结构体，是handler层传给service的，注意handelr和service交互用submittasjkreq
	//和taskres这两个结构体，而和repo层交互一律使用task这个结构体。这样分层的作用是数据就够也一起解耦

	if req.Name == "" {
		return nil, errors.NewUserError(4001, "task name can not be empty", nil)
	}

	ctx := context.Background()
	s.queuerepo.Enqueue(ctx, id, req.DelayTime)
	s.queuerepo.SetStatusCache(ctx, id, "initial")

	return &model.TaskRes{
		Status: "submitted",
		ID:     id,
	}, nil
}

// 获取任务状态的接口，返回给前端，前端根据状态来轮询，直到变成done状态
func (s *TaskService) GetTaskStatus(id int64) (*model.TaskRes, error) {
	if id <= 0 {
		return nil, errors.NewUserError(4002, "task id invalid", nil)
	}

	ctx := context.Background()
	status, err := s.queuerepo.GetStatusCache(ctx, id)
	if err == nil && status != "" {
		return &model.TaskRes{
			ID:     id,
			Status: status,
		}, nil
	}

	task, err := s.repo.GetTaskStatus(id)
	if err != nil {

		//区分是不是没有这一行数据的错误，否则就返回system error
		if err == sql.ErrNoRows {
			return nil, errors.NewUserError(4041, "task not found", err)
		}
		return nil, errors.NewSystemError(5002, "[Service Error]failed to get the taskid", err)
	}
	//交给handler这个taskres就是用来和handler层交互的
	return &model.TaskRes{
		ID:     task.ID,
		Status: task.Status,
		Result: task.Result,
	}, nil
}
