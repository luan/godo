package godo

import (
	"errors"
	"github.com/coopernurse/gorp"
	"strconv"
)

type Task struct {
	Id     int
	Name   string
	Status string
}

func (t *Task) NextStatus() string {
	if t.Status == "done" {
		return "pending"
	}
	return "done"
}

var Dbmap *gorp.DbMap

func AddTask(name string) Task {
	t := Task{Name: name, Status: "pending"}
	_ = Dbmap.Insert(&t)
	return t
}

func FindTask(id int) (t Task, err error) {
	obj, err := Dbmap.Get(Task{}, id)
	if err == nil && obj != nil {
		t = *(obj.(*Task))
	} else if err == nil {
		err = errors.New("Couldn't find task with Id: " + strconv.Itoa(id))
	}

	return
}

func Tasks() (tasks []Task) {
	_, _ = Dbmap.Select(&tasks, "select * from tasks order by id")
	return
}

func UpdateTask(task Task) (_, err error) {
	_, err = Dbmap.Update(&task)
	return
}

func ResetTasks() {
	Dbmap.TruncateTables()
}
