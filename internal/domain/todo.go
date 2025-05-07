package domain

type Todo struct {
	ID          uint
	Title       string
	Description string
	Status      Status
}

func (Todo) TableName() string {
	return "todos"
}
