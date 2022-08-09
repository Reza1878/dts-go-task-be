package task

type CreateTaskRequest struct {
	Task     string `json:"task" binding:"required"`
	Assignee string `json:"assignee" binding:"required"`
	DueDate  string `json:"due_date" binding:"required"`
}

type GetTaskRequest struct {
	Id int `uri:"id" binding:"required"`
}
