package task

type Service interface {
	CreateTask(request CreateTaskRequest) (Task, error)
	GetTask(request GetTaskRequest) (Task, error)
	GetTasks() ([]Task, error)
	UpdateTask(taskId GetTaskRequest, request CreateTaskRequest) (Task, error)
	DeleteTask(taskId GetTaskRequest) error
}
