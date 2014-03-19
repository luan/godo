package godo

import (
	"errors"
	"strconv"
)

type Task struct {
	Id     int
	Name   string `form:"name" binding:"required"`
	Status string `form:"status"`
}

func (t Task) NextStatus() string {
	if t.Status == "done" {
		return "pending"
	}
	return "done"
}

var numberOfTasks int
var tasks []Task

func AddTask(name string) Task {
	numberOfTasks++
	t := Task{Name: name, Status: "pending", Id: numberOfTasks}
	tasks = append(tasks, t)
	return t
}

func FindTask(id int) (t *Task, err error) {
	for i := range tasks {
		if tasks[i].Id == id {
			return &tasks[i], nil
		}
	}

	err = errors.New("Couldn't find task with Id: " + strconv.Itoa(id))
	return
}

func Tasks() []Task {
	return tasks
}

func ResetTasks() {
	tasks = []Task{}
	numberOfTasks = 0
}
