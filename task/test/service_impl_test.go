package task

import (
	"dts-task/app"
	"dts-task/task"
	"testing"
)

func TestGetTasks(t *testing.T) {
	repository := task.NewRepository(app.NewDB())

	service := task.NewService(repository)

	tasks, err := service.GetTasks()

	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	for _, task := range tasks {
		t.Log(task.Task)
		t.Log(task.Id)
	}
}

func TestGetTask(t *testing.T) {
	repository := task.NewRepository(app.NewDB())

	service := task.NewService(repository)

	_, err := service.GetTask(task.GetTaskRequest{Id: 8})

	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
}

func TestCreateTask(t *testing.T) {
	repository := task.NewRepository(app.NewDB())

	service := task.NewService(repository)

	task, err := service.CreateTask(task.CreateTaskRequest{
		Task:     "Mengerjakan DTS Task",
		Assignee: "Me",
		DueDate:  "2022-08-09",
	})

	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	t.Log(task)
}

func TestUpdateTask(t *testing.T) {
	repository := task.NewRepository(app.NewDB())

	service := task.NewService(repository)

	task, err := service.UpdateTask(
		task.GetTaskRequest{Id: 8},
		task.CreateTaskRequest{
			Task:     "Mengerjakan DTS Task Update",
			Assignee: "Me",
			DueDate:  "2022-08-15",
		})

	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	t.Log(task)
}

func TestDeleteTask(t *testing.T) {
	repository := task.NewRepository(app.NewDB())

	service := task.NewService(repository)

	err := service.DeleteTask(
		task.GetTaskRequest{Id: 7})

	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	t.Log(err)
}
