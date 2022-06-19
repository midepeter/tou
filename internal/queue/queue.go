package queue

type Queue struct{
	ID  int
	Name string
	Store string
}

func NewQueue(id int, name, store string) *Queue {
	return &Queue{
		ID: id,
		Name: name,
		Store: store,
	}
}