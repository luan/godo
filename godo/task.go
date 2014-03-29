package godo

type Task struct {
	ID     int
	Name   string
	Status string
}


func (t *Task) NextStatus() string {
	if t.Status == "done" {
		return "pending"
	}
	return "done"
}

