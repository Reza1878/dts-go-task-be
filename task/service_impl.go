package task

import (
	"dts-task/helper"
)

type serviceImpl struct {
	repository Repository
}

func NewService(repository Repository) *serviceImpl {
	return &serviceImpl{
		repository: repository,
	}
}

func (s *serviceImpl) GetTasks() ([]Task, error) {
	var tasks []Task

	tasks, err := s.repository.FindAll()

	return tasks, err
}

func (s *serviceImpl) GetTask(request GetTaskRequest) (Task, error) {
	task, err := s.repository.FindById(request.Id)

	return task, err
}

func (s *serviceImpl) CreateTask(request CreateTaskRequest) (Task, error) {
	task := Task{}
	time, err := helper.ParseDate(request.DueDate)
	if err != nil {
		return task, err
	}

	task.Assignee = request.Assignee
	task.DueDate = time
	task.Task = request.Task
	task, err = s.repository.Save(task)

	return task, err
}

func (s *serviceImpl) UpdateTask(requestId GetTaskRequest, request CreateTaskRequest) (Task, error) {
	_, err := s.repository.FindById(requestId.Id)
	if err != nil {
		return Task{}, err
	}

	time, err := helper.ParseDate(request.DueDate)
	if err != nil {
		return Task{}, err
	}
	task := Task{
		Assignee: request.Assignee,
		DueDate:  time,
		Task:     request.Task,
		Id:       requestId.Id,
	}

	task, err = s.repository.Update(task)

	return task, err
}

func (s *serviceImpl) DeleteTask(request GetTaskRequest) error {
	_, err := s.repository.FindById(request.Id)
	if err != nil {
		return err
	}

	err = s.repository.Delete(request.Id)
	return err
}
