package model

import "time"

//repo层的数据模型
type Task struct {
	ID        int64
	Name      string
	Status    string
	Result    string
	DelayTime int
}

//service层的请求和响应模型

type TaskSubmit struct {
	Name      string `json:"name"`
	DelayTime int    `json:"delay_time"`
}

type TaskRes struct {
	ID     int64  `json:"task_id"`
	Status string `json:"status"`
	Result string `json:"result,omitempty"`
}

type TaskRedis struct {
	Name      string `json:"name"`
	ID        int64  `json:"task_id"`
	DelayTime int    `json:"delay_time"`
}

type DistRes struct {
	Allow   bool
	Current int64
	TTL     time.Duration
}
