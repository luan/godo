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

func (tm *TaskManager) FindTasksOfProject(projectId int, tasks *[]Task) (err error) {
	err = Dbmap.Select(tasks, "select * from tasks where projectID = ? order by id", projectId)
	return
}
