package service

import (
	"context"
	"database/sql"
	"sync"

	"myapp/internal/app/repo"
	"myapp/internal/model"
	"myapp/pkg/errors"
	"myapp/pkg/snowid"
)

type TaskService struct {
	repo      *repo.TaskRepo
	redisrepo *repo.RedisRepo

	workerpool         int64
	jobqueue           chan *model.TaskRedis
	processconcurrency chan struct{}
	wg                 sync.WaitGroup
}

// 在service层把这些配置注入
func NewTaskService(ra *repo.TaskRepo, rb *repo.RedisRepo, workpool int64, jobqueue int64, processcurrency int64) *TaskService {
	return &TaskService{
		repo:      ra,
		redisrepo: rb,

		workerpool:         workpool,
		jobqueue:           make(chan *model.TaskRedis, jobqueue),
		processconcurrency: make(chan struct{}, processcurrency),
	}
}

// 提交任务的接口，接受handler层传来的submittaskreq结构体，返回taskres结构体
func (s *TaskService) SubmitTask(ctx context.Context, req model.TaskSubmit) (*model.TaskRes, error) {

	//生成一个全局唯一的id，作为任务id，后续根据这个id来查询任务状态和结果
	i, err := snowid.NextID()
	//雪花id的格式不是int64，所以需要转换一下，转换失败就返回system error
	id := int64(i)
	if err != nil {
		return nil, errors.NewSystemError(5003, "snowid error", err)
	}

	//这里接受的是一个submittaskreq结构体，是handler层传给service的，注意handelr和service交互用submittasjkreq
	//和taskres这两个结构体，而和repo层交互一律使用task这个结构体。这样分层的作用是数据就够也一起解耦

	if req.Name == "" {
		return nil, errors.NewUserError(4001, "task name can not be empty", nil)
	}

	//入队，入队失败就返回system error
	err = s.redisrepo.Enqueue(ctx, id, req.Name, req.DelayTime)
	if err != nil {
		return nil, errors.NewSystemError(5001, "failed to enqueue the task", err)
	}
	//把任务状态设置成queued写入redis缓存
	s.redisrepo.SetStatusCache(ctx, id, "queued")

	return &model.TaskRes{
		Status: "submitted",
		ID:     id,
	}, nil
}

// 获取任务状态的接口，返回给前端，前端根据状态来轮询，直到变成done状态
func (s *TaskService) GetTaskStatus(ctx context.Context, id int64) (*model.TaskRes, error) {
	if id <= 0 {
		return nil, errors.NewUserError(4002, "task id invalid", nil)
	}

	//热点数据优化，缓解数据库的压力，先从redis缓存中查询，如果有就直接返回，如果没有再去数据库查询
	status, err := s.redisrepo.GetStatusCache(ctx, id)
	if err == nil && status != "" {
		return &model.TaskRes{
			ID:     id,
			Status: status,
		}, nil
	}

	task, err := s.repo.GetTaskStatus(ctx, id)
	if err != nil {

		//区分是不是没有这一行数据的错误，否则就返回system error
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError(4041, "task not found", err)
		}
		return nil, errors.NewSystemError(5002, "[Service Error]failed to get the taskid", err)
	}
	s.redisrepo.SetStatusCache(ctx, id, task.Status)

	//交给handler这个taskres就是用来和handler层交互的
	return &model.TaskRes{
		ID:     task.ID,
		Status: task.Status,
		Result: task.Result,
	}, nil
}
