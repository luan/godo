package godo

type Task struct {
	ID     int
	Name   string
	Status string
	ProjectID int
}


type TaskDomain struct {
	Name   string
	Status string
}

func (t *Task) NextStatus() string {
	if t.Status == "done" {
		return "pending"
	}
	return "done"
}

func NewTask(name string) Task {
	return Task{Name: name, Status: "pending"}
}

func NewTaskWithProjectID(name string, project_id int) Task {

	return Task{Name: name, ProjectID: project_id, Status: "pending"}
}

func NewTaskDomain(name string, status string) Task {
	return Task{Name: name, Status: status}
}



