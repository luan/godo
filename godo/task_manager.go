package godo

type TaskManager struct {
	manager
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (tm *TaskManager) Add(name string) (t Task, err error) {
	t = Task{Name: name, Status: "pending"}
	err = tm.manager.Add(&t)
	return
}

func (tm *TaskManager) FindAll() (tasks []Task, err error) {
	err = tm.manager.FindAll(&tasks, "tasks")
	return
}
