package repo

import (
	"app/internal/model"
	"database/sql"
)

// 解耦用于DI注入
type TaskRepo struct {
	DB *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{
		DB: db,
	}
}

//repo层具有的方法，负责有关数据库的操作，这里不能返回errors.NewxxxError

func (r *TaskRepo) CreateTask(t *model.Task) (int64, error) {
	res, err := r.DB.Exec("INSERT INTO tasks (name,status,delay_time) VALUES (?,'pending',?)", t.Name, t.DelayTime)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *TaskRepo) UpdateStatus(id int64, status string) error {
	_, err := r.DB.Exec("UPDATE tasks SET status=? WHERE id = ?", status, id)
	return err
}

func (r *TaskRepo) GetTaskStatus(id int64) (*model.Task, error) {
	var t model.Task
	var status, result sql.NullString

	err := r.DB.QueryRow("SELECT status,result FROM tasks WHERE id = ?", id).Scan(&status, &result)
	if err != nil {
		return nil, err
	}

	t.ID = id
	t.Status = status.String
	t.Result = result.String

	return &t, nil
}
