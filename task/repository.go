package task

type Repository interface {
	Save(task Task) (Task, error)
	FindAll() ([]Task, error)
	FindById(Id int) (Task, error)
	Update(task Task) (Task, error)
	Delete(Id int) error
}
