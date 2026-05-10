package repo

import (
	"context"
	"database/sql"
	"myapp/internal/model"
)

// TaskRepo结构体
type TaskRepo struct {
	DB *sql.DB
}

// 创建TaskRepo实例
func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{
		DB: db,
	}
}

//创建任务

func (r *TaskRepo) CreateTask(ctx context.Context, id int64, t *model.Task) (int64, error) {
	res, err := r.DB.ExecContext(ctx, "INSERT INTO tasks (id,name,status,delay_time) VALUES (?,?,?,?)", id, t.Name, t.Status, t.DelayTime)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// 更新任务状态
func (r *TaskRepo) UpdateStatus(ctx context.Context, id int64, status string) error {
	_, err := r.DB.ExecContext(ctx, "UPDATE tasks SET status=? , updated_at=CURRENT_TIMESTAMP WHERE id = ?", status, id)
	return err
}

// 获取任务状态
func (r *TaskRepo) GetTaskStatus(ctx context.Context, id int64) (*model.Task, error) {
	var t model.Task
	var status, result sql.NullString
	//Replace with id parameter
	err := r.DB.QueryRowContext(ctx, "SELECT status,result FROM tasks WHERE id = ?", id).Scan(&status, &result)
	if err != nil {
		return nil, err
	}

	t.ID = id
	t.Status = status.String
	t.Result = result.String

	return &t, nil
}
