package task

import (
	"dts-task/app"
	"dts-task/task"
	"fmt"
	"testing"
	"time"
)

func TestTaskInsert(t *testing.T) {
	taskRepository := task.NewRepository(app.NewDB())

	time, err := time.Parse("2006-01-02", "2022-08-09")

	if err != nil {
		t.Error("Failed to parse date")
	}

	task := task.Task{
		Task:     "Second task",
		Assignee: "John",
		DueDate:  time,
	}

	entity, err := taskRepository.Save(task)

	if err != nil {
		fmt.Println(err)
		t.Error("Failed to create task")
	}

	fmt.Println(entity.Id)
}

func TestFindById(t *testing.T) {
	r := task.NewRepository(app.NewDB())

	task, err := r.FindById(2)

	if err != nil {
		t.Errorf("Failed to get data, error: %s", err.Error())
	}

	t.Log(task)
}

func TestUpdate(t *testing.T) {
	r := task.NewRepository(app.NewDB())

	time, err := time.Parse("2006-01-02", "2022-08-09")

	if err != nil {
		t.Error("Failed to parse date")
	}

	task := task.Task{
		Id:       2,
		Task:     "Second task update",
		Assignee: "John Doe",
		DueDate:  time,
	}

	newTask, err := r.Update(task)

	if err != nil {
		t.Fatalf("Failed to update data, error: %s", err.Error())
	}

	t.Log(newTask)
}

func TestDelete(t *testing.T) {
	r := task.NewRepository(app.NewDB())

	err := r.Delete(2)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
}
