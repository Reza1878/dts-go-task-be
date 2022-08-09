package task

import (
	"database/sql"
)

type repositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repositoryImpl {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) Save(task Task) (Task, error) {

	stmt, err := r.db.Prepare("INSERT INTO tasks (task, assignee, due_date) VALUES (?, ?, ?)")

	if err != nil {
		return task, err
	}

	result, err := stmt.Exec(task.Task, task.Assignee, task.DueDate)

	if err != nil {
		return task, err
	}

	id, err := result.LastInsertId()

	task.Id = int(id)

	return task, err
}

func (r *repositoryImpl) FindAll() ([]Task, error) {

	var tasks []Task
	rows, err := r.db.Query("SELECT id, task, assignee, due_date from tasks")
	if err != nil {
		return tasks, err
	}

	for rows.Next() {
		task := Task{}
		rows.Scan(&task.Id, &task.Task, &task.Assignee, &task.DueDate)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *repositoryImpl) FindById(Id int) (Task, error) {
	task := Task{}

	script := "SELECT id, task, assignee, due_date FROM tasks WHERE id = ?"
	rows := r.db.QueryRow(script, Id)

	err := rows.Scan(&task.Id, &task.Task, &task.Assignee, &task.DueDate)

	return task, err
}

func (r *repositoryImpl) Update(task Task) (Task, error) {

	script := "UPDATE tasks SET task = ?, assignee = ?, due_date = ? WHERE id = ?"

	_, err := r.db.Exec(script, task.Task, task.Assignee, task.DueDate, task.Id)

	return task, err
}

func (r *repositoryImpl) Delete(Id int) error {

	script := "DELETE FROM tasks WHERE id = ?"

	_, err := r.db.Exec(script, Id)

	return err
}
