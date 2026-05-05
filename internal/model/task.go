package model

//repo层的数据模型
type Task struct {
	ID        int64
	Name      string
	Status    string
	Result    string
	DelayTime int
}

//service层的请求和响应模型

type SubmitTaskReq struct {
	Name      string `json:"name"`
	DelayTime int    `json:"delay_time"`
}

type TaskRes struct {
	TaskID int64  `json:"task_id"`
	Status string `json:"status"`
	Result string `json:"result,omitempty"`
}
