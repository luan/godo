package godo

type TaskManager struct {
	manager
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (tm *TaskManager) FindAll() (tasks []Task, err error) {
	err = tm.manager.FindAll(&tasks, "tasks")
	return
}
