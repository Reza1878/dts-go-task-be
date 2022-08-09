package task

import "time"

type Task struct {
	Id       int       `json:"id"`
	Task     string    `json:"task"`
	Assignee string    `json:"assignee"`
	DueDate  time.Time `json:"due_date"`
}
